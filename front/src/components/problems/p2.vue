<template>
  <div class="problem2">
    <h1>Hello {{ name }}!!</h1>
    <p>
      {{ statement }}
      <br>
      <b>Password: P455W0RD</b>
      <br>
      <input type="text" placeholder="passwordはここに入力!" v-model="pass" maxlength=4>
      <button @click="getFlag">Get the flag!</button>
    </p>

    <p>
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
  name: 'P2',
  data () {
    return {
      name: firebase.auth().currentUser.email.split('@')[0],
      statement: '下にパスワードがあるから, パスワードを入力して答えをGETしてね!'
    }
  },
  created () {
    const p2 = firebase.database().ref(`all/${this.name}/p2`)
    p2.update({
      flag: '-'
    })
  },
  methods: {
    getFlag: function () {
      if (this.pass === 'P455W0RD') {
        Axios.post('http://192.168.33.10:8080/api/v1/problems/getFlag', {
          'id': 1,
          'pass': decodeURIComponent(escape(atob('UzFMVkVSQlVMTEVU')))
        })
          .then(response => {
            console.log(response)
            const res = response.data.flag
            if (res) {
              alert('Flag is 「' + res + '」!')
            } else {
              alert('The answer is incorrect...')
            }
          })
          .catch(err => {
            alert(err.response)
            console.log(err.response)
          })
      }
    },
    submit: function () {
      Axios.post('http://192.168.33.10:8080/api/v1/problems/check', {
        id: 2,
        ans: this.flag
      })
        .then(response => {
          console.log(response)
          const res = response.data.result
          if (res) {
            alert('Congratulations!\nYou\'ve got ' + response.data.score + ' Points!')
            this.$router.push('/ranking')
            const p2 = firebase.database().ref(`all/${this.name}/p2`)
            p2.update({
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
