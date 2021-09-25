
import axios from 'axios'

const client = axios.create({
  baseURL: process.env.NOTIF_ENDPOINT,
  // withCredentials: true,
  timeout: 120000
})

export default client
