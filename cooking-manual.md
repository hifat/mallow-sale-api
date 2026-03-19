## Google Client ID

---

### ขั้นตอน

**1. สร้าง Project**
- ไปที่ [console.cloud.google.com](https://console.cloud.google.com)
- คลิก **Select a project** → **New Project** → ตั้งชื่อ → Create

**2. เปิดใช้ Google API**
- ไปที่ **APIs & Services** → **Library**
- ค้นหา **Google Drive API** → Enable
- ค้นหา **Google Identity** (หรือ People API) → Enable

**3. สร้าง OAuth Consent Screen**
- ไปที่ **APIs & Services** → **OAuth consent screen**
- **Branding** → ใส่ App name + email
- **Audience** → ตรงนี้แหละที่เลือก External/Internal

**4. สร้าง Credentials**
- ไปที่ **APIs & Services** → **Credentials**
- **Create Credentials** → **OAuth 2.0 Client ID**
- Application type: **Web application**
- Authorized JavaScript origins:
  ```
  http://localhost:<your-port>     ← dev
  https://yourdomain.com    ← production
  ```
- Authorized redirect URIs: (ถ้าใช้ FE-only ไม่ต้องใส่)
- กด **Create** → จะได้ **Client ID** มา

---

### หน้าตา Client ID

```
4xxxxxxxxx-abcdefghijk.apps.googleusercontent.com
```

---

### เอาไปใส่

```env
# FE (.env)
VITE_GOOGLE_CLIENT_ID=4xxxxxxxxx-abcdefghijk.apps.googleusercontent.com

# BE (.env)
GOOGLE_CLIENT_ID=4xxxxxxxxx-abcdefghijk.apps.googleusercontent.com
```

---

> ⚠️ **หมายเหตุ:** ตอน deploy production อย่าลืมกลับมาเพิ่ม domain จริงใน **Authorized JavaScript origins** ด้วยไม่งั้น Google จะ block request
