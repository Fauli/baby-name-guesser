<script setup lang="ts">
import { ref, watchEffect } from 'vue'

const API_URL = `http://localhost:8080/api/votes/top`

const names: any = ref<string[] | null>(null)

var voteStatus = ref('')
var failureReason = ref('')

var counter = 1

watchEffect(() => {
  voteStatus.value = ''
  failureReason.value = ''
})

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}`
  names.value = await (await fetch(url)).json()
})


const selectedNames = ref<string[]>([])

function handleNameSelection(name: string) {
  if (selectedNames.value.includes(name)) {
    selectedNames.value = selectedNames.value.filter(n => n !== name)
  } else {
    selectedNames.value.push(name)
  }
}

function submitVote() {
  console.log('Submitting vote for', selectedNames.value)
  var values = selectedNames.value
  fetch('http://localhost:8080/api/votes/top', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      names: values
    })
  })
  .then(response => response.json())
  .then(data => {
    console.log('Success:', data);
    // if errorcode is not 200, show an alert
    if (data.email.length = 0)
    {
      voteStatus.value = 'failure'
      failureReason.value = data.message
    } else {
      voteStatus.value = 'success'
    }
    
  })
  .catch((error) => {
    console.error('Error:', error);
    alert('Error submitting vote')
  });
}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else-if="names.message">
    <p>Cannot fetch the name list: {{ names.message }}</p>
  </div>
  <div v-else>
    <h1>👑 Top 3 guesses 👑</h1>
    <p>Get a sneak peek on the 3 top voted names so far!</p>
    <hr>
    <br>
    <ul>
      <p v-for="(key, item) in names" :key="item">
        <span class="message"> &nbsp; {{ counter++ }}. {{ key }}</span><br>
      </p>
    </ul>

    <br>
    <hr>
    <br>

    <div v-if="voteStatus === 'success' || voteStatus === 'failure'">
      <div class="important">
        <div v-if="voteStatus === 'success'" class="important-content green">
          <p>Voting successful!
            <br />Now you can <RouterLink to="/voting">Vote!</RouterLink>
          </p>
        </div>
        <div v-if="voteStatus === 'failure'" class="important-content red">
          <!-- TODO: [franz] Add reason why voting has failed here! -->
          <p>Voting has failed!</p>
          <p>Reason: {{ failureReason }}</p>
        </div>
      </div>
    </div>

    <div v-if="selectedNames.length > 0">
      <p>You will be voting for:</p>
      <ul>
        <li v-for="name in selectedNames" :key="name">{{ name }}</li>
      </ul>
      <p>This results in a bet of <span style="font-weight: bold; text-decoration: underline;">{{ selectedNames.length * 10 }} Franken</span></p>
      <p>You can use twint, revolut or paypal - please mark your name when transferring 💙</p>
    </div>
    <br>
    <button @click="submitVote" class="submit-button">Submit Vote</button>


  </div>
</template>

<style>

.important {
  padding: 10px;
  border-radius: 5px;
  margin-bottom: 20px;
}

.red {
  background-color: #f8d7da;
  border: 1px solid #f5c6cb;
  border-radius: 5px;
  margin-bottom: 20px;
  padding: 10px;

}

.green {
  background-color: #d4edda;
  border: 1px solid #c3e6cb;
  border-radius: 5px;
  margin-bottom: 20px;
  padding: 10px;

}

hr {
  border: none;
  border-top: 1px solid #ccc;
  margin: 20px 0;
}

.submit-button {
  background-color: #42b883;
  color: white;
  padding: 10px 20px;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.submit-button:hover {
  background-color: #32a672;
}

a {
  text-decoration: none;
  /* color: #42b883; */
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