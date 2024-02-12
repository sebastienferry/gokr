<template>
    <div>{{name}}</div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
    data() {
        return {
            name: 'About'
        };
    },
})
</script>