/// <reference lib="WebWorker" />

// export empty type because of tsc --isolatedModules flag
export type {}
declare const self: ServiceWorkerGlobalScope


import { initializeApp } from "firebase/app"
import { getMessaging, onBackgroundMessage } from "firebase/messaging/sw"
import { INotif } from "./models/notif"

// Initialize the Firebase app in the service worker by passing in
// your app's Firebase config object.
// https://firebase.google.com/docs/web/setup#config-object

const firebaseApp = initializeApp({
  apiKey: process.env.API_KEY,
  authDomain: process.env.AUTH_DOMAIN,
  projectId: process.env.PROJECT_ID,
  storageBucket: process.env.STORAGE_BUCKET,
  messagingSenderId: process.env.MSGID,
  appId: process.env.APP_ID,
  measurementId: process.env.MEASURE_ID
})

// Retrieve an instance of Firebase Messaging so that it can handle background
// messages.
const messaging = getMessaging(firebaseApp)

onBackgroundMessage(messaging, (payload) => {
  const notif: INotif = payload.data as any
  self.registration.showNotification(notif.message)
  console.log(payload)
})