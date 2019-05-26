<template>
  <div class="signin">
    <h2>Sign in</h2>
    <input type="text" placeholder="username" v-model="username">
    <input type="password" placeholder="password" v-model="password">
    <button @click="signin">Sign in</button>
    <p>
      Don't you have an account?
      <router-link to="/signup">create your account now</router-link>
    </p>
  </div>
</template>

<script>
import firebase from 'firebase'

export default {
  name: 'Signin',
  data () {
    return {
      username: '',
      password: ''
    }
  },
  methods: {
    signin: function () {
      console.log(process.env.VUE_APP_APIKEY)
      firebase
        .auth()
        .signInWithEmailAndPassword(this.username, this.password)
        .then(
          user => {
            alert('Success!')
            this.$router.push('/')
          },
          err => {
            alert(err.message)
          }
        )
    }
  }
}
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
.signup {
  margin-top: 20px;

  display: flex;
  flex-flow: column nowrap;
  justify-content: center;
  align-items: center;
}
input {
  margin: 10px 0;
  padding: 10px;
}
</style>
