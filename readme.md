# 🍪 CJE - Cookie Jar Exploiter

Go-based security tool designed to test web applications for vulnerabilities related to cookie handling.
---

## 🔍 Features

- ✅ Reads a list of subdomains from a file
- ✅ Checks each domain for:
  - **Cookie Jar Overflow**: Tests if a server accepts too many cookies
  - **Session Fixation**: Checks if the server accepts a user-defined session ID
- ✅ ANSI-colored CLI output for better visibility
- ✅ Efficient, fast HTTP requests with timeout handling

---

## ⚠️ Why Use This Tool?

- **Cookie Jar Overflow** can cause browser instability and potential DoS in client-side environments.
- **Session Fixation** is a serious vulnerability that can allow attackers to hijack valid user sessions.

---

## 🛠️ Installation

```bash
git clone https://github.com/progprnv/CJE.git
cd CJE
go build -o CJE main.go
