<template>
    <div>{{name}}</div>    
</template>

<script lang="ts">
import { defineComponent, computed } from 'vue';
import { useStore } from 'vuex'
export default defineComponent({
    data() {
        const store = useStore()
        return {
            name: 'Home',
            isAuthenticated: computed(() => store.state.isAuthenticated)
        };
    },
})
</script>