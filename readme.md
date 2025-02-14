# Base Golang Raiden + Supabase Project

## 🚀 Overview

This project is built using Golang Raiden and Supabase to create a full-stack backend service with authentication, database, and real-time functionalities.

## 🛠 Tech Stack

- **Backend:** Golang Raiden Framework
- **Database:** Supabase (PostgreSQL)
- **Authentication:** Supabase Auth
- **Storage:** Supabase Storage
- **Deployment:** Docker, Railway, or VPS

## 📌 Features

- User Authentication (Sign up, Sign in, JWT-based auth)
- RESTful API with Golang Raiden
- Realtime Data Updates
- Secure Role-Based Access
- File Storage with Supabase Storage

## 🏗 Installation & Setup

**1️⃣ Clone the Repository**

```sh
git clone https://github.com/yassar021/be-raiden.git
cd your-repo
```

**2️⃣ Install Dependencies**

```sh
go mod tidy
```

**3️⃣ Set Up Environment Variables**

```sh
SUPABASE_URL=your_supabase_url
SUPABASE_ANON_KEY=your_anon_key
SUPABASE_SERVICE_ROLE_KEY=your_service_role_key
PORT=8080
```

**4️⃣ Run the Development Server**

```sh
go run cmd/medpoint.go
```

This will launch the app on the connected emulator or device.

## 🚀 Deployment

**Build and Release**

- Android
  ```sh
  flutter build apk
  flutter build appbundle
  ```
- iOS
  `sh
    flutter build ios
    `
  Make sure to set up Xcode and an Apple Developer Account.

<!-- ## 📂 Project Structure
```bash
├── lib/
│   ├── main.dart
│   ├── screens/
│   ├── components/
│   ├── providers/
│   ├── services/
│   ├── utils/
├── assets/               # Static assets (images, fonts, etc.)
├── pubspec.yaml          # Flutter dependencies
└── README.md             # Project documentation -->

```

## 🙌 Acknowledgments
- [Golang Raiden](https://raiden.sev-2.com/docs/quick-start)
- [Supabase](https://supabase.com/)


```
