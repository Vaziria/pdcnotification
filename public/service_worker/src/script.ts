import { setupNotification } from "./notification"

const envType = process.env.PROJECT_ID

export default class Main {
  constructor() {
    if ("serviceWorker" in navigator) {
      navigator.serviceWorker
        .register("/firebase-messaging-sw.js", { scope: "/" })
        .then(function () {
          console.log("Service Worker Registered")
        })
    }


    this.initializeMessaging()
  }

  initializeMessaging(): void {
    const email: string | undefined = (window as any).pdc_email
    if(email) {
      setupNotification(email, (data) => {
        console.log("messaging received data", data)

      }).then(()=>{
        console.log(`email ${email} setup notification finish`)
      })
    }
  }

  testprint(): void {
    const doc = document.getElementById('title')
    doc.innerText = `PROJECT ID : ${envType}`
  }
}

new Main()