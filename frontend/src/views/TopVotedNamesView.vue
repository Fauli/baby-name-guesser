<script setup lang="ts">
import { ref, watchEffect } from 'vue'

const API_URL = `/api/votes/top`

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
  fetch('/api/votes/top', {
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

function getMedal(counter: number) {
  if (counter == 1) {
    return "🥇"
  } else if (counter == 2) {
    return "🥈"
  } else if (counter == 3) {
    return "🥉"
  } else {
    return "🏅"
  }
}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else-if="names.message">
    <h1>Not quite yet...</h1>
    <p>You will only be able to view the top 3 guessed names so far after you guessed yourself ⭐️</p>
    <br />
    <p style="font-style: italic;">Details: {{ names.message }}</p>
  </div>
  <div v-else>
    <h1>Top 3 guesses 📊</h1>
    <p>Get a sneak peek on the 3 top guessed names so far...</p>
    <hr>
    <ul style="padding-left: 0px;">
      <p v-for="(key, item) in names" :key="item">
        <span style="font-size: xx-large;" class="message"> &nbsp; {{ getMedal(counter++)}} {{ key }}</span><br>
      </p>
    </ul>


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


</style>