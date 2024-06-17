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

const email = ref('')
const password = ref('')

var loginStatus = ref('')
var failureReason = ref('')

watchEffect(() => {
  loginStatus.value = ''
})

function login() {
  const voter = {
    email: email.value,
    password: password.value
  }

  fetch('/api/voters/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(voter)
  })
    .then(response => response.json())
    .then(data => {
      console.log('Success:', data);
      // if something failed, no mail is returned
      if (!data.email) {
        console.log('Login failed')
        loginStatus.value = 'failure'
        failureReason.value = data.message
      } else {
        console.log('Login successful')
        loginStatus.value = 'success'
        router
          .push({ path: '/voting' })
          .then(() => { router.go(0) })
      }

    })
    .catch((error) => {
      console.error('Error:', error);
      loginStatus.value = 'failure'
      failureReason.value = error
      alert('Error submitting vote: ' + error)
    });
}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else>
    <h1>Login for the name game!</h1>
    <p>Enter your email and password to login, if you do not have a username and password yet, register first</p>

    <div v-if="loginStatus === 'success' || loginStatus === 'failure'">
      <div class="important">
        <div v-if="loginStatus === 'success'" class="important-content green">
          <p>Login successful!
            <br />Now you can <RouterLink to="/voting">Vote!</RouterLink>
          </p>
        </div>
        <div v-if="loginStatus === 'failure'" class="important-content red">
          <p>Login failed!</p>
          <p>Details: {{ failureReason }}</p>
        </div>
      </div>
    </div>
  </div>
  <div>


    <form>
      <label for="email">Email:</label>
      <input type="email" id="email" v-model="email" required>

      <label for="password">Password:</label>
      <input type="password" id="password" v-model="password" required>

      <button type="submit" @click="login">Login</button>
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
