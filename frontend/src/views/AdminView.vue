<script setup lang="ts">
import { ref, watchEffect, computed } from 'vue';

const API_URL = `/api`;

const voters: any = ref('');
const votes: any = ref('');
const votedNames: any = ref<string[]>([]);
const sortKey = ref('');
const sortOrder = ref(1); // 1 for ascending, -1 for descending

watchEffect(async () => {
  const url = `${API_URL}/voters`;
  voters.value = await (await fetch(url)).json();
});

watchEffect(async () => {
  const url = `${API_URL}/votes/voters`;
  votedNames.value = await (await fetch(url)).json();
});

watchEffect(async () => {
  const url = `${API_URL}/votes`;
  votes.value = await (await fetch(url)).json();
});

function userPaid(email: string) {
  console.log('Marking user as paid:', email);
  fetch('/api/payments/' + email, {
    method: 'POST',
  })
    .then((response) => response.json())
    .then((data) => {
      console.log('Success:', data);
      if (!data.message.includes('has been processed')) {
        console.log('Marking as paid failed');
        alert('Marking as paid failed: ' + data.message);
      } else {
        console.log('Marking as paid successful');
        voters.value = voters.value.map((voter: any) => {
          if (voter.email === email) {
            voter.has_paid = true;
          }
          return voter;
        });
      }
    })
    .catch((error) => {
      console.error('Error:', error);
    });
}

function sortBy(key: string) {
  if (sortKey.value === key) {
    sortOrder.value *= -1;
  } else {
    sortKey.value = key;
    sortOrder.value = 1;
  }
  voters.value = [...voters.value].sort((a: any, b: any) => {
    if (a[key] > b[key]) return sortOrder.value;
    if (a[key] < b[key]) return -sortOrder.value;
    return 0;
  });
}

const sortedVotes = computed(() => {
  return Object.entries(votes.value)
    .sort(([, aVotes], [, bVotes]) => (bVotes as number) - (aVotes as number))
    .map(([name, voteCount]) => ({ name, voteCount }));
});

const totalCollectedMoney = computed(() => {
  return Object.values(votes.value).reduce((acc, voteCount) => acc as number + (voteCount as number) * 10, 0);
});
</script>

<template>
  <div v-if="!voters">
    <p>You probably don't have access to this resource :)</p>
  </div>
  <div v-else-if="voters.message">
    <p>{{ voters.message }}</p>
  </div>
  <div v-else>
    <h1>Admin Panel</h1>
    <hr />

    <h2>Users</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th @click="sortBy('has_voted')">Voted</th>
            <th @click="sortBy('has_paid')">Paid</th>
            <th @click="sortBy('name')">Name</th>
            <th @click="sortBy('last_name')">Last Name</th>
            <th @click="sortBy('email')">Email</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="voter in voters" :key="voter.id">
            <td>
              <span v-if="voter.has_voted">✅</span><span v-else>❌</span>
            </td>
            <td>
              <a @click="userPaid(voter.email)">
                <span v-if="voter.has_paid">💰</span><span v-else>❌</span>
              </a>
            </td>
            <td>{{ voter.name }}</td>
            <td>{{ voter.last_name }}</td>
            <td>{{ voter.email }}</td>
          </tr>
        </tbody>
      </table>
    </div>

    <br /><br />
    <h2>Votes by User</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Voter</th>
            <th>Names</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="votedName in votedNames">
            <td>{{ votedName.email }}</td>
            <td v-if="votedName.names">
              <span v-for="name in votedName.names">{{ name }}, </span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>

    <br /><br />

    <h2>Votes</h2>
    <div class="table-container">
      <table>
        <thead>
          <tr>
            <th>Name</th>
            <th>Votes</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="vote in sortedVotes" :key="vote.name">
            <td>{{ vote.name }}</td>
            <td>{{ vote.voteCount }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>

  <br /><br />

  <!-- can be calculated by summing up all value os votes and multiply by 10-->
  <h2>Collected money</h2>
  <p>{{ totalCollectedMoney }} CHF</p>

</template>

<style>
body {
  font-family: 'Arial', sans-serif;
  background-color: #f4f4f9;
  margin: 0;
  padding: 20px;
}

h1 {
  color: #333;
}

h2 {
  color: #555;
  margin-bottom: 10px;
}

.table-container {
  width: 100%;
  overflow-x: auto;
  background-color: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

thead {
  background-color: #42b883;
  color: white;
  cursor: pointer;
}

th, td {
  padding: 12px;
  text-align: left;
}

th {
  font-weight: bold;
  user-select: none;
}

tr:nth-child(even) {
  background-color: #f9f9f9;
}

a {
  color: #42b883;
  cursor: pointer;
}

hr {
  border: none;
  border-top: 1px solid #ccc;
  margin: 20px 0;
}
</style>
