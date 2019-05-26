import Vue from 'vue'
import Router from 'vue-router'
import Home from '@/components/Home'
import Score from '@/components/Score'
import Signup from '@/components/Signup'
import Signin from '@/components/Signin'
import P1 from '@/components/problems/p1'
import P2 from '@/components/problems/p2'
import firebase from 'firebase'

Vue.use(Router)

let router = new Router({
  routes: [
    {
      path: '*',
      redirect: 'signin'
    },
    {
      path: '/',
      name: 'Home',
      component: Home,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/signup',
      name: 'Signup',
      component: Signup
    },
    {
      path: '/signin',
      name: 'Signin',
      component: Signin
    },
    {
      path: '/ranking',
      name: 'Ranking',
      component: Score
    },
    {
      path: '/p1',
      name: 'P1',
      component: P1,
      meta: {
        requireAuth: true
      }
    },
    {
      path: '/p2',
      name: 'P2',
      component: P2,
      meta: {
        requireAuth: true
      }
    }
  ]
})

router.beforeEach((to, from, next) => {
  let requireAuth = to.matched.some(record => record.meta.requireAuth)
  if (requireAuth) {
    firebase.auth().onAuthStateChanged(function (user) {
      if (user) {
        next()
      } else {
        next({
          path: '/signin',
          query: {
            redirect: to.fullPath
          }
        })
      }
    })
  } else {
    next()
  }
})

export default router
