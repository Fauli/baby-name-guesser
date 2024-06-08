<script setup lang="ts">
import { ref, watchEffect } from 'vue'
import { useRoute } from 'vue-router'

/*
{
    "email_address": "string",
    "event_password": "string",
    "id": "string",
    "invited_by": "string",
    "name": "string",
    "status": "string",
    "subscribed_on": "string",
    "tags": [
      "string"
    ]
  }

  */

interface Attendee {
    email_address: string;
    event_password: string;
    id: string;
    invited_by: string;
    name: string;
    status: string;
    subscribed_on: string;
    tags: string[];
}

const API_URL = `https://qreator.ch/api/attendees/`

const attendee: any = ref(null)
const errorMessage: any = ref(null)
const successMessage: any = ref(null)
const route = useRoute()

// TODO: get attendeeId from route
// const attendeeId = "0472fbece42e54e47b2ab15f1b8ee663"
const attendeeId = route.params.id

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  // const url = `${API_URL}${attendeeId}`
  // var id = this.$route.params.id
  const url = `${API_URL}${attendeeId}`
  attendee.value = await (await fetch(url)).json().catch((error) => {
    console.error('Error:', error);
    // alert('Error getting attendee ' + attendeeId + ' Error: ' + error)
    errorMessage.value = error.message
    attendee.value = "{}"
  })
})

function getCheckInPath(item: Attendee) {
  // path is https://qreator.ch/attendees/$id/check-ins
  var root = 'https://qreator.ch/api/attendees/'
  var id = item.id
  var checkInPath = root.concat(id, '/check-ins')
  console.log(checkInPath)

  return checkInPath
}

// function checkInAttendee(item) performs a POST call to the check-in endpoint
function checkInAttendee(item: Attendee) {
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
  .then(response => {
    console.log(response)
    if (!response.id) {
             // create error object and reject if not a 2xx response code
             let err = new Error(response.message)
             //err.response = response
             //err.status = response.status
             throw err
         } else {
          successMessage.value = "Checked in!"
          
         }
  })
  .then(data => {
    console.log('Success:', data);
  })
  .catch((error) => {
    console.error('Error:', error);
    // alert('Error checking in ' + item.value + ' Error: ' + error)
    errorMessage.value = error.message
  });
} 

// function setTagColor sets the color of the text in the class to red if the value is "no" and green if the value is "yes"
function setTagColor(tag: string) {
  if (tag == "helfer") {
    return "red"
  } else if (tag == "gaesteliste") {
    return "green"
  } else {
    return "black"
  }
}

</script>

<template>
    <div v-if="errorMessage">
    <h1 style="color: red">Error!</h1>
    <p>{{ errorMessage }}</p>
  </div>
  <div v-if="successMessage">
    <h1 style="color: green">Success!</h1>
    <p>{{ successMessage }}</p>
  </div>
  <div v-if="$route.params.id == 'null'">
    <p>
      No id provided, please provide one :)
    </p>

  </div>
  <div v-if="!attendee">
    <h1>Loading...</h1>
    <p>Getting information for id {{ $route.params.id }}</p>
  </div>
  <div v-else v-bind="attendee">
      <h1>{{ attendee.name }}</h1><br>
      <a @click="checkInAttendee(attendee)">Check in {{ attendee.name }}</a><br><br>
      <span class="message">E-Mail: {{ attendee.email_address }}</span><br>
      <!-- <span class="message">{{ attendee.id }}</span><br> -->
      <span class="message">Eingeladen von: {{ attendee.invited_by }}</span><br>
      <p>Tags:</p>
      <div v-for="tag in attendee.tags" :key="tag">
        <span :class="setTagColor(tag)">{{ tag }}</span>
      </div><br>
  </div>
</template>

<style>
a {
  text-decoration: none;
  color: #42b883;
}
li {
  line-height: 1.5em;
  margin-bottom: 20px;
}
.author,
.date {
  font-weight: bold;
}

.red {
  color: red;
  font-weight: bold;
}

.green {
  color: green;
  font-weight: bold;
}
</style>