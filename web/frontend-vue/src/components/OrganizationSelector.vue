<template>
    <div>
        <select class="select" onchange="">
            <option v-for="organization in organizations" :key="organization.id" :value="organization.id">{{ organization.name }}</option>
        </select>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
    name: 'OrganizationSelector',
    setup() {
        return {
            organizations: [
                {
                    id: 1,
                    name: 'Organization 1'
                },
                {
                    id: 2,
                    name: 'Organization 2'
                },
            ]    
        }
    },
    mounted() {
        console.log('OrganizationSelector mounted');
    }
});
</script>
