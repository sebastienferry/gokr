<template>
  <nav class="level">
    <!-- Left side -->
    <div class="level-left"> 
      <div class="level-item">
        <p class="title is-3">
        <div>🎯Gokr</div>
        </p>
      </div>
      <div class="level-item">
        <div class="">
          <nav class="navbar" role="navigation" aria-label="main navigation">
            <div class="navbar-brand">
              <a role="button" class="navbar-burger" aria-label="menu" aria-expanded="false"
                data-target="navbarBasicExample">
                <span aria-hidden="true"></span>
                <span aria-hidden="true"></span>
              </a>
            </div>

            <div id="navbar" class="navbar-menu">
              <div class="navbar-start">
                <router-link class="navbar-item" to="/ui">Home</router-link>                
                <div class="navbar-item has-dropdown is-hoverable">                  
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
          </nav>
        </div>
      </div>
    </div>
    <div class="level-right">
      <div class="buttons">
        <!-- <a class="button is-primary">
          <strong>Sign up</strong>
        </a> -->
        <OrganizationSelector></OrganizationSelector>
        <a class="button is-light">
          Log Out
        </a>
      </div>
    </div>
  </nav>
  <div class="container">
    <router-view />
  </div>  
</template>

<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
  name: 'App',
  data: () => ({
    isAuthenticated: false,
  }),
});
</script>