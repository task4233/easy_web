<template>
  <div class="problem1">
    <h1>Hello {{ name }}!!</h1>
    <p>
      {{ statement }}
      <br>
      <b>FLAG{TH15_15_54MPLE_FL4G}</b>
      <br>
      <input type="text" placeholder="FLAG{...}" v-model="flag">
      <button @click="submit">submit</button>
    </p>
    <router-link to="/">topに戻る</router-link>
    <button @click="signout">Sign Out</button>
  </div>
</template>

<script>
import Vue from 'vue'
import Axios from 'axios'
import VueAxios from 'vue-axios'
import firebase from 'firebase'

Vue.use(VueAxios, Axios)

export default {
  name: 'P1',
  data () {
    return {
      name: firebase.auth().currentUser.email.split('@')[0],
      statement: 'この問題はサンプル問題です。'
    }
  },
  created () {
    const p1 = firebase.database().ref(`all/${this.name}/p1`)
    p1.update({
      flag: '-'
    })
  },
  methods: {
    submit: function () {
      Axios.post('http://192.168.33.10:8080/api/v1/problems/check', {
        id: 1,
        ans: this.flag
      })
        .then(response => {
          console.log(response)
          const res = response.data.result
          if (res) {
            alert('Correct!\nYou\'ve got ' + response.data.score + ' Points!')
            this.$router.push('/ranking')
            const p1 = firebase.database().ref(`all/${this.name}/p1`)
            p1.update({
              flag: 'o'
            })
          } else {
            alert('The answer is incorrect...')
          }
        })
        .catch(err => {
          console.log(err.response)
        })
    },
    signout: function () {
      firebase
        .auth()
        .signOut()
        .then(() => {
          alert('Sign out complete!')
          this.$router.push('/signin')
        })
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
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
</style>
