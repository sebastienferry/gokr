import { createApp } from 'vue'
import { createStore } from 'vuex'
import App from './App.vue'
import './style.css'
import router from './router'

// Create a new store instance.
const store = createStore({
    state () {
      return {
        isAuthenticated: false,
      }
    },
    mutations: {      
      authenticate (state) {        
        state.isAuthenticated = true
      },
      logout (state) {
        state.isAuthenticated = false
      }
    }
  })

// Create the app instance.
const app = createApp(App)
app.use(store)
app.use(router)
app.mount('#app')


