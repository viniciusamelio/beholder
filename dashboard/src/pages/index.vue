<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
import { getEnvironmentsQuery } from '@/composables/queries/getEnvironmentsQuery';
import { getSessionsQuery } from '@/composables/queries/getSessionsQuery';
import type { Environment } from '@/types';
import { computed, ref, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
const { data: environments, isPending } = getEnvironmentsQuery();
const selectedEnvironment = ref<Environment | null>(null);
const route = useRoute()
const router = useRouter()
const id = computed(() => selectedEnvironment.value?.id);
const sessionsQuery = getSessionsQuery({ id: Number(route.params.id) });

watch(selectedEnvironment, async (environment) => {
  if (!environment) return;
  router.push({ name: "/", params: { id: environment.id } })
  await sessionsQuery.refetch();
})
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
