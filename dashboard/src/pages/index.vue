<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
import { getEnvironmentsQuery } from '@/composables/queries/getEnvironmentsQuery';
import type { Environment } from '@/types';
import { ref } from 'vue';
const { data: environments, isPending } = getEnvironmentsQuery();
const selectedEnvironment = ref<Environment | null>(null);
</script>

<template>
  <section class="grid grid-cols-2 gap-4">
    <article class="px-6 py-4">
      <div v-if="isPending" class="flex flex-col gap-2">
        <div v-for="i in 8" class="rounded-lg p-4 w-full h-32 bg-card animate-pulse" />
      </div>
      <div v-else class="flex flex-col gap-2">
        <EnvironmentCard @click="selectedEnvironment = environment" v-for="environment in environments"
          :environment="environment" :selected="environment.id == selectedEnvironment?.id" />
      </div>
    </article>

  </section>
</template>
