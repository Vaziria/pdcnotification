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


    // this.initializeMessaging()
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

  initializeNotification(email: string): Promise<void> {
    if(email) {
      return setupNotification(email, (data) => {
        console.log("messaging received data", data)

      }).then(()=>{
        console.log(`email ${email} setup notification finish`)
      })
    }
  }

}

(window as any)['pdcnotification'] = new Main()