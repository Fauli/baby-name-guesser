<script setup lang="ts">
import { ref, watchEffect } from 'vue'

const API_URL = `http://localhost:8080/api/names`

const names: any = ref<string[] | null>(null)


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
  fetch('http://localhost:8080/api/votes', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      email: 'f@sbebe.ch',
      names: selectedNames.value
    })
  })
  .then(response => response.json())
  .then(data => {
    console.log('Success:', data);
    alert('Vote submitted')
    // if errorcode is not 200, show an alert
    if (data.status !== 200) {
      alert('Error: ' + data.error);
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
  <div v-else>
    <h1>List of all possible names</h1>
    <p>This site shows all discussed names.</p>
    <p>There are currently <b>{{ names.names.length }}</b> available to choose from.</p>
   
    <br>
    <hr>
    <br>
    <ul>
      <p v-for="item in names.names" :key="item.value">
      <input type="checkbox" v-model="selectedNames" :value="item" />
      <span class="message"> &nbsp; {{ item }}</span><br>
      </p>
    </ul>

    <br>
    <hr>
    <br>

    <p v-if="selectedNames.length > 0">You will be voting for:</p>
    <ul>
      <li v-for="name in selectedNames" :key="name">{{ name }}</li>
    </ul>
    <br>
    <button @click="submitVote" class="submit-button">Submit Vote</button>




  </div>
</template>

<style>

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