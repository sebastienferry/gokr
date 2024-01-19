<script lang="ts">
import { defineComponent } from 'vue';
import OrganizationSelector from './components/OrganizationSelector.vue';
export default defineComponent({
  name: 'App',
  components: {
    OrganizationSelector,
  },
})
</script>

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
                <a class="navbar-item">
                  Home
                </a>

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
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
