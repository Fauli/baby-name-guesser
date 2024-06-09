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

const email = ref('')
const password = ref('')

var loginStatus = ref('')

watchEffect(() => {
  loginStatus.value = ''
})

const register = async () => {
  const voter = {
    email: email.value,
    password: password.value
  }

  try {
    const response = await fetch('http://localhost:8080/api/voters/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(voter)
    })



    if (response.ok) {
      // Registration successful
      console.log('Login successful')
      loginStatus.value = 'success'
    } else {
      // Registration failed
      console.error('Login failed')
      loginStatus.value = 'failure'
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
    <h1>Login for the name game!</h1>

    <div v-if="loginStatus === 'success' || loginStatus === 'failure'">
      <div class="important">
        <div v-if="loginStatus === 'success'" class="important-content green">
          <p>Login successful!
            <br />Now you can <RouterLink to="/voting">Vote!</RouterLink>
          </p>
        </div>
        <div v-if="loginStatus === 'failure'" class="important-content red">
          <p>Login failed!</p>
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

      <button @click="register">Register</button>
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
