<template>
  <div class="Ranking">
    <h1>{{ msg }}</h1>
    <button @click="signout">Sign Out</button>
  </div>
</template>

<script>
import firebase from 'firebase'

export default {
  name: 'Score',
  data () {
    return {
      msg: 'Welcome to Score Page(Alpha)!',
      name: firebase.auth().currentUser.email.split('@')[0],
    }
  },
  methods: {
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
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
</style>
