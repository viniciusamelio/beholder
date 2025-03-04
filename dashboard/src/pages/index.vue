<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
import { useGetRequestsQuery } from '@/composables/queries/getRequestsQuery';
import { useGetSessionsQuery } from '@/composables/queries/getSessionsQuery';
import { useEnvironments } from '@/composables/useEnvironments';

const environment = useEnvironments();
const sessionsQuery = useGetSessionsQuery();
const requestsQuery = useGetRequestsQuery();
const { isPending, data: envs } = environment.environments();

environment.$subscribe((_, state) => {
  if (state.selectedEnvironment) {
    sessionsQuery.refetch(state.selectedEnvironment.id);
    requestsQuery.refetch(state.selectedEnvironment.id);
  }
});
</script>

<template>
  <section class="grid grid-cols-2 gap-4">
    <article class="px-6 py-4">
      <div v-if="isPending" class="flex flex-col gap-2">
        <div v-for="i in 8" class="rounded-lg p-4 w-full h-32 bg-card animate-pulse" />
      </div>
      <div v-else class="flex flex-col gap-2">
        <EnvironmentCard @click="environment.selectEnvironment(env)" v-for="env in envs"
          :environment="env" :selected="environment.selectedEnvironment?.id == env.id" />
      </div>
    </article>

  </section>
</template>
