<script setup lang="ts">
import router from '@/router';
import { ref, watchEffect } from 'vue'

const API_URL = `/api/names`

const names: any = ref<string[] | null>(null)


watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}`
  names.value = await (await fetch(url)).json()
})

const name = ref('')
const lastname = ref('')
const email = ref('')
const password = ref('')
const event_password = ref('')

var registrationStatus = ref('')
var registrationMessage = ref('')

watchEffect(() => {
  registrationStatus.value = ''
})

function register() {
  console.log('Registering voter...')
  const voter = {
    name: name.value,
    last_name: lastname.value,
    email: email.value,
    password: password.value,
    event_password: event_password.value
  }

  fetch('/api/voters', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(voter)
  })
    .then(response => response.json())
    .then(data => {
      if (data.email) {
        // Registration successful
        console.log('Registration successful')
        registrationStatus.value = 'success'
        router
          .push({ path: '/voting' })
          .then(() => { router.go(0) })

      } else {
        // Registration failed
        console.error('Registration failed')
        registrationStatus.value = 'failure'
        registrationMessage.value = 'Registration failed, please validate if all information are correct. Details: ' + data.message
      }


    })
    .catch((error) => {
      console.error('Error:', error);
      registrationStatus.value = 'failure'
      registrationMessage.value = error
      alert('Error submitting vote: ' + error)
    });

}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else>
    <h1>Register for the name game!</h1>
    <p>You need the special password (Event password) in order to register, ask Franz or Kathrin for it :).</p>
    <br />
    <br />

  <div v-if="registrationStatus === 'success' || registrationStatus === 'failure'">
      <div class="important">
        <div v-if="registrationStatus === 'success'" class="important-content green">
          <p>
            Registration successful!
          </p>
        </div>
        <div v-if="registrationStatus === 'failure'" class="important-content red">
          <p>Registration failed!</p>
        </div>
      </div>
    </div>


    <form>
      <label for="name">Your Name:</label>
      <input type="text" id="name" v-model="name" required>
      
      <label for="lastname">Your Lastname:</label>
      <input type="text" id="lastname" v-model="lastname" required>
      
      <label for="email">Your Email (will also be your username):</label>
      <input type="email" id="email" v-model="email" required>
      
      <label for="password">Your Password:</label>
      <input type="password" id="password" v-model="password" required minlength="8">

      <label for="password">Event Password:</label>
      <input type="text" id="event_password" v-model="event_password" required>
      
      <button type="submit" @click="register">Register</button>
    </form>
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
.important-content {
  font-weight: bold;
}


a {
  text-decoration: none;
  /* color: #42b883; */
}

.author,
.date {
  font-weight: bold;
}


a {
  text-decoration: none;
  /* color: #42b883; */
}
.author,
.date {
  font-weight: bold;
}

form {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 10px;
  max-width: 400px;
  margin: 0 auto;
  padding: 20px;
  border: 1px solid #ccc;
  border-radius: 5px;
  background-color: #f5f5f5;
}

label {
  font-weight: bold;
}

input[type="text"],
input[type="email"],
input[type="password"] {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
}

button {
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  background-color: #42b883;
  color: white;
  cursor: pointer;
}

button:hover {
  background-color: #3da566;
}

</style>
