import { getMessaging, getToken, onMessage } from "firebase/messaging"
import { initializeApp } from "firebase/app"
import api from './api/client'
import { INotif } from "./models/notif"

const firebaseApp = initializeApp({
    apiKey: process.env.API_KEY,
    authDomain: process.env.AUTH_DOMAIN,
    projectId: process.env.PROJECT_ID,
    storageBucket: process.env.STORAGE_BUCKET,
    messagingSenderId: process.env.MSGID,
    appId: process.env.APP_ID,
    measurementId: process.env.MEASURE_ID
})

const messaging = getMessaging(firebaseApp)

export type Callback = (notif: INotif) => void

export async function installToken(email: string, tokens: string): Promise<void> {
    const notif: INotif = {
        email,
        message: '',
        action: 'add_token',
        tokens: [tokens]
    }
    const res = await api.post('/Notification', notif)
}


export async function setupNotification (email: string, callback: Callback): Promise<void> {
  try {
    const token = await getToken(messaging, {
        vapidKey: process.env.VAPIKEY
    })

    try {
      await installToken(email, token)
      onMessage(messaging, (notif) => callback(notif.data as any))
    } catch (e) {
      console.error(e)
    }
  } catch (e) {
    console.error(e)
  }
}
