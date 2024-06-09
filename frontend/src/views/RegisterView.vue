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

const name = ref('')
const lastname = ref('')
const email = ref('')
const password = ref('')
const event_password = ref('')

var registrationStatus = ref('')

watchEffect(() => {
  registrationStatus.value = ''
})

const register = async () => {
  const voter = {
    name: name.value,
    last_name: lastname.value,
    email: email.value,
    password: password.value,
    event_password: event_password.value
  }

  try {
    const response = await fetch('http://localhost:8080/api/voters', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(voter)
    })



    if (response.ok) {
      // Registration successful
      console.log('Registration successful')
      registrationStatus.value = 'success'
    } else {
      // Registration failed
      console.error('Registration failed')
      registrationStatus.value = 'failure'
    }
  } catch (error) {
    console.error('Error:', error)
  }
}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else>
    <h1>Register for the name game!</h1>
    <p>You need the special password in order to register, ask Franz or Kathrin for it :).</p>
    <br />
    <br />

    <div v-if="registrationStatus === 'success' || registrationStatus === 'failure'">
    <div class="important">
      <div class="important-content">
        <p v-if="registrationStatus === 'success'">Registration successful!
          <br />Now you can <RouterLink to="/voting">Vote!</RouterLink>
</p>
        <p v-if="registrationStatus === 'failure'">Registration failed!</p>
        
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
      <input type="password" id="password" v-model="password" required>

      <label for="password">Event Password:</label>
      <input type="text" id="event_password" v-model="event_password" required>
      
      <button @click="register">Register</button>
    </form>
  </div>


</template>

<style>

.important {
  background-color: #f8d7da;
  color: #721c24;
  padding: 10px;
  border: 1px solid #f5c6cb;
  border-radius: 5px;
  margin-bottom: 20px;
}

.important-content {
  font-weight: bold;
}


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
