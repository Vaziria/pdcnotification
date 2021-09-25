# pdcnotification

# endpoint

untuk endpoint production tanya ke developer.
untuk running lokal
`go run cmd/main.go`

# payload endpoint
```
{
	action: 'send_action' | 'add_token'
	email: string
	message: string
	token: string[]
}
```

# snippet install pligin
```
<script>
  var exports = {"__esModule": true}
</script>
<script type="module" src="/script.js"></script>
<script>
window.pdcnotification.initializeNotification('user@gmail.com')
</script>
```


# deploy tutorial

- gcloud functions deploy Notification --runtime go116 --trigger-http --allow-unauthenticated



# localhost firebase
```
<script type="module">
  // Import the functions you need from the SDKs you need
  import { initializeApp } from "https://www.gstatic.com/firebasejs/9.0.2/firebase-app.js";
  import { getAnalytics } from "https://www.gstatic.com/firebasejs/9.0.2/firebase-analytics.js";
  // TODO: Add SDKs for Firebase products that you want to use
  // https://firebase.google.com/docs/web/setup#available-libraries

  // Your web app's Firebase configuration
  // For Firebase JS SDK v7.20.0 and later, measurementId is optional
  const firebaseConfig = {
    apiKey: "AIzaSyB-bvRwNZAWRCKEFVTxE-gErCE0Kg2YmX8",
    authDomain: "pdc-base.firebaseapp.com",
    projectId: "pdc-base",
    storageBucket: "pdc-base.appspot.com",
    messagingSenderId: "671716308705",
    appId: "1:671716308705:web:3a39142eb8eba6712ba9a0",
    measurementId: "G-71XRDCZDYE"
  };

  // Initialize Firebase
  const app = initializeApp(firebaseConfig);
  const analytics = getAnalytics(app);
</script>
```