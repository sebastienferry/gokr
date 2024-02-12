<template>
  <nav class="navbar" role="navigation" aria-label="main navigation">
    
    
    <div class="navbar-brand">
      <div class="navbar-item">
        <p class="is-large subtitle">
        <div>🎯Gokr</div>
        </p>
      </div>
      <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false" data-target="navbarBasicExample">
        <span aria-hidden="true"></span>
        <span aria-hidden="true"></span>
      </a>
    </div>
    
    
    <div id="navbar" class="navbar-menu">
      <div class="navbar-start">
        <router-link class="navbar-item" to="/ui" v-if="isAuthenticated">Home</router-link>
        <div class="navbar-item has-dropdown is-hoverable" v-if="isAuthenticated">
          <a class="navbar-link">
            Configuration
          </a>
          <div class="navbar-dropdown">
            <a class="navbar-item">
              Organizations
            </a>
          </div>
        </div>
        <router-link class="navbar-item" to="/ui/about">About</router-link>        
      </div>
    </div>

    
    <div class="navbar-end">      
      <div class="navbar-item">
        <div class="buttons" v-if="!isAuthenticated">
          <button class="button is-primary" @click="authenticate">
            <strong>Login</strong>
          </button>
        </div>
      </div>
      
      <div class="navbar-item">
        <div class="buttons" v-if="isAuthenticated">
          <div>
            <OrganizationSelector />
          </div>
          <button class="button is-light" @click="logout">
            Log Out
          </button>
        </div>
      </div>
    </div>
  </nav>

  <div class="container">
    <router-view />
  </div>
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'vuex'
import OrganizationSelector from './components/OrganizationSelector.vue';

export default defineComponent({
  name: 'App',
  components: {
    OrganizationSelector
},
  data() {
    const store = useStore()
    return {
      isAuthenticated: computed(() => store.state.isAuthenticated),
      authenticate: () => { store.commit('authenticate')},
      logout: () => store.commit('logout')
    }
  }
});
</script>