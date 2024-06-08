<script setup lang="ts">
import { ref, watchEffect } from 'vue'

const API_URL = `https://qreator.ch/api/qrcodes`

const qrCodes: any = ref<QRCode[] | null>(null)

interface QRCode {
    id: string;
    name: string;
    attendeeName: string;
}

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}`
  qrCodes.value = await (await fetch(url)).json()
})

function getImgPath(item: QRCode) {
  // https://storage.googleapis.com/qreator-qrcodes/output/03ff45487ac1f88c72f1172e4cc22703.png
  return `https://storage.googleapis.com/qreator-qrcodes/output/${item.id}.png`
}

function getCheckInPath(item: QRCode) {
  // path is https://qreator.ch/attendees/$id/check-ins
  var root = 'https://qreator.ch/api/attendees/'
  var id = item.name.split('.')[0]
  var checkInPath = root.concat(id, '/check-ins')

  return checkInPath
}

// function checkInAttendee(item) performs a POST call to the check-in endpoint
function checkInAttendee(item: QRCode) {
  var checkInPath = getCheckInPath(item)
  fetch(checkInPath, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    // body: JSON.stringify({
    //   attendeeId: item.name.split('.')[0]
    // })
  })
  .then(response => response.json())
  .then(data => {
    console.log('Success:', data);
    alert('Checked in ' + item.name)
  })
  .catch((error) => {
    console.error('Error:', error);
    alert('Error checking in ' + item.name + ' Error: ' + error)
  });
} 

</script>


<template>
  <div v-if="!qrCodes">
    <h1>Loading...</h1>
  </div>
  <div v-else>
    <h1>List of all QR codes</h1>
    <p>This site shows all QR codes generated for the event.</p>
    <p>There are currently <b>{{ qrCodes.length }}</b> on the server.</p>
    <br>
    <hr>
    <br>
    <ul>
      <li v-for="item in qrCodes" :key="item.value">
        <span class="message">{{item.attendeeName}}</span><br>
          <img class="qr-code" :src="getImgPath(item)"><br>
          <a @click="checkInAttendee(item)">Check in {{item.attendeeName}}</a>
      </li>
    </ul>
  </div>
</template>

<style>
a {
  text-decoration: none;
  /* color: #42b883; */
}
li {
  line-height: 1.5em;
  margin-bottom: 20px;
}
.author,
.date {
  font-weight: bold;
}

.qr-code {
  width: 256px;
  height: 256px;
}

</style>