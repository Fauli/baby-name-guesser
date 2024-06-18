<script setup lang="ts">
import { ref, watchEffect } from 'vue'

const API_URL = `/api`

const names: any = ref<string[] | null>(null)

var voteStatus = ref('')
var failureReason = ref('')

const userInfo: any = ref('')
const paymentInfo: any = ref('')
const votedNames: any = ref<string[]>([])

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}/me`
  userInfo.value = await (await fetch(url)).json()
})

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}/votes/me`
  votedNames.value = await (await fetch(url)).json()
})

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}/payments`
  paymentInfo.value = await (await fetch(url)).json()
})

watchEffect(() => {
  voteStatus.value = ''
  failureReason.value = ''
})

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}/names`
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
  fetch('/api/votes', {
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
    if (!data.email)
    {
      voteStatus.value = 'failure'
      failureReason.value = data.message
    } else {
      voteStatus.value = 'success'
      setTimeout(() => {
        location.reload()
      }, 2000)
    }
    
  })
  .catch((error) => {
    console.error('Error:', error);
    voteStatus.value = 'failure'
    failureReason.value = error
    alert('Error submitting vote: ' + error)
  });
}

</script>


<template>
  <div v-if="!names">
    <h1>Loading...</h1>
  </div>
  <div v-else-if="names.message">
    <h1>Register and Login to bet 🥳</h1>
    <p>After you are logged-in you can place your bets!</p>
    <br />
    <p style="font-style: italic;">Details: {{ names.message }}</p>
  </div>
  <div v-else>
    <div v-if="userInfo.has_voted">
      <h1>Thank you for voting 💙</h1>
      <br />
      <p>Transfer the bet money to one of the following accounts to finalize your participation:</p>
      <table>
        <tr>
          <td class="label">Amount</td>
          <td class="value">{{ votedNames.length * 10  }}.00</td>
        </tr>
        <tr>
          <td class="label">Twint</td>
          <td class="value">{{ paymentInfo.twint }}</td>
        </tr>
        <tr>
          <td class="label">Revolut</td>
          <td class="value">{{ paymentInfo.revolut }}</td>
        </tr>
        <tr>
          <td class="label">Paypal</td>
          <td class="value">{{ paymentInfo.paypal }}</td>
        </tr>
        <tr>
          <td class="label">IBAN</td>
          <td class="value">{{ paymentInfo.iban }}</td>
        </tr>        <tr>
          <td class="label">Important</td>
          <td class="value">Add your name to the payment</td>
        </tr>
      </table>
      <br />
      <p v-if="!userInfo.has_paid">👉 <span style="text-decoration: underline;">Sometime after you have transferred the money, you will see a confirmation message on this site.</span> 👈</p>
      <p v-else>🎉<span class="rainbow-text" style="font-size: larger; font-weight: bold; text-decoration: underline;">Your money was recieved and you are officially taking part in the game! Woohoo!</span>🎉</p>
      <br />
      <p>Thank you for participating in the name game!</p>
      <p>We will be sending out the results after the 20.08.2024 🥰</p>
      <br />
      <h2>Your votes:</h2>
      <ul>
        <li v-for="name in votedNames" :key="name">{{ name }}</li>
      </ul>
    </div>
    <div v-else>
      <h1>Place your guess 🍼</h1>
      <p>Select all your guesses, but choose wisely. You can only guess once!</p>
      <p>There are <b>{{ names.names.length }}</b> available names to choose from.</p>
      <p>The vote only really counts once you transferred the matching sum.</p>
      <p>You will get the payment information after you placed your votes.</p>
      <hr>
      <ul>
        <p v-for="item in names.names" :key="item.value">
          <input type="checkbox" v-model="selectedNames" :value="item" />
          <span class="message"> &nbsp; {{ item }}</span><br>
        </p>
      </ul>

      <br>
      <hr>
      <br>

      <div v-if="voteStatus === 'success' || voteStatus === 'failure'">
        <div class="important">
          <div v-if="voteStatus === 'success'" class="important-content green">
            <p>Voting successful!</p>
            <p>You will be redirected in a few seconds...</p>
          </div>
          <div v-if="voteStatus === 'failure'" class="important-content red">
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
        <p>This results in a bet of <span style="font-weight: bold; text-decoration: underline;">{{ selectedNames.length
            * 10 }} Franken</span></p>
        <p>You can use twint, revolut or paypal - please mark your name when transferring 💙</p>
      </div>
      <br>
      <button @click="submitVote" class="submit-button">Submit Vote</button>
    </div>



  </div>
</template>

<style>

.label {
  width: 100px;
}

.value {
  font-weight: bold;
}

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


.rainbow-text {
  background-image: linear-gradient(to left, violet, indigo, blue, green, orange, red);
  -webkit-background-clip: text;
  -moz-background-clip: text;
  background-clip: text;
  -webkit-text-fill-color: transparent;
}


</style>