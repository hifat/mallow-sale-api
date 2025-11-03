package shoppingRepository

import (
	"context"
	"fmt"
	"io"

	billingReaderPb "github.com/hifat/billing-reader/pb"
	shoppingModule "github.com/hifat/mallow-sale-api/internal/shopping"
	"google.golang.org/grpc"
)

type receiptGRPCRepository struct {
	grpcConn *grpc.ClientConn
}

func NewReceiptGRPC(grpcConn *grpc.ClientConn) IReceiptGrpcRepository {
	return &receiptGRPCRepository{grpcConn}
}

func (r *receiptGRPCRepository) ReadReceipt(ctx context.Context, fileName string, file []byte) ([]shoppingModule.ResReceiptReader, error) {
	c := billingReaderPb.NewBillingReaderClient(r.grpcConn)
	readRcpStream, err := c.ReadReceipt(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to read receipt: %v", err)
	}

	chunkSize := 5 * 1024 * 1024

	chunkCount := 0
	for i := 0; i < len(file); i += chunkSize {
		end := i + chunkSize
		if end > len(file) {
			end = len(file)
		}

		chunk := file[i:end]
		chunkCount++

		err := readRcpStream.Send(&billingReaderPb.ReadReceiptRequest{
			FileName: fileName,
			Chunk:    chunk,
		})
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("Failed to send file: %v", err)
		}

		// TODO: Remove this after develop
		fmt.Printf("Sent chunk %d: %d bytes\n", chunkCount, len(chunk))
	}

	fmt.Printf("Streaming completed. Total chunks: %d\n\n", chunkCount)

	fmt.Println("waiting for response...")
	rcpRes, err := readRcpStream.CloseAndRecv()
	if err != nil {
		return nil, fmt.Errorf("Failed to CloseAndRecv: %v", err)
	}

	if !rcpRes.Success {
		return nil, fmt.Errorf("Response error: %v", rcpRes.Error)
	}

	fmt.Printf("✅ Receipt processed successfully\n")

	res := make([]shoppingModule.ResReceiptReader, 0, len(rcpRes.Receipt))
	for _, v := range rcpRes.Receipt {
		if v != nil {
			res = append(res, shoppingModule.ResReceiptReader{
				InventoryID:      "",
				Name:             v.Name,
				NameEdited:       v.NameEdited,
				PurchasePrice:    v.PurchasePrice,
				PurchaseQuantity: v.PurchaseQuantity,
				Remark:           v.Remark,
			})
		}
	}

	return res, nil
}
