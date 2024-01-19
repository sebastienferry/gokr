<template>
    <div class="select">        
        <select id="organization" v-model="selectedOrganization">
            <option v-for="organization in organizations" :key="organization.id" :value="organization.id">
                {{ organization.name }}
            </option>
        </select>
    </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
export default defineComponent({
    data() {
        return {
            selectedOrganization: null,
            organizations: [
                { id: 1, name: 'Organisation 1' },
                { id: 2, name: 'Organisation 2' },
                { id: 3, name: 'Organisation 3' },
            ],
        };
    }
})
</script>
