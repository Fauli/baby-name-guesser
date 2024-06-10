<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import { ref, watchEffect } from 'vue'

const API_URL = `/api`

const userInfo: any = ref('')

// TODO: Maybe check https://auth0.com/blog/beginner-vuejs-tutorial-with-user-login/

watchEffect(async () => {
  // this effect will run immediately and then
  // re-run whenever currentBranch.value changes
  const url = `${API_URL}/me`
  userInfo.value = await (await fetch(url)).json()
})
</script>


<template>
  <header>
    <div class="wrapper">
      <nav>
      <img alt="Vue logo" class="logo" src="@/assets/logo.png" width="75" height="75" />
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/voting">Vote!</RouterLink>
        <!-- <RouterLink to="/check-in/null">Check-In</RouterLink> -->
        <span class="navItem" v-if="userInfo.name">
          <RouterLink to="/top-votes">See top 3 votes</RouterLink>
          <span style="border-left: 1px solid var(--color-border); padding: 0 1rem;">Logged in as <span style="font-weight: bold;">{{ userInfo.name }} {{ userInfo.last_name }}</span>
          </span>
          <!-- <a href="/api/voters/logout">Logout</a> -->
        </span>
        <span class="navItem" v-else>
          <RouterLink to="/login">Login</RouterLink>
          <RouterLink to="/register">Register</RouterLink>
        </span>
      </nav>
      </div>
  </header>

  <RouterView />
</template>

<style scoped>

img:hover {
  /* Start the shake animation and make the animation last for 0.5 seconds */
  animation: shake 0.5s;

  /* When the animation is finished, start again */
  animation-iteration-count: infinite;
}

@keyframes shake {
  0% { transform: translate(0, 0) rotate(0deg); }
  25% { transform: translate(5px, 5px) rotate(5deg); }
  50% { transform: translate(0, 0) rotate(0eg); }
  75% { transform: translate(-5px, 5px) rotate(-5deg); }
  100% { transform: translate(0, 0) rotate(0deg); }}


nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  /* color: var(--color-text); */
  color: #000000;
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

.navItem {
  display: inline-block;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  header {
    display: flex;
    place-items: center;
    padding-right: calc(var(--section-gap) / 2);
  }

  .logo {
    margin: 0 2rem 0 0;
  }

  header .wrapper {
    display: flex;
    place-items: flex-start;
    flex-wrap: wrap;
    width: 100%;

  }

  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>