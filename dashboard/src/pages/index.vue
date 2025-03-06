<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
import RequestCard from '@/components/molecules/RequestCard.vue';
import SessionCard from '@/components/molecules/SessionCard.vue';
import RequestTabSkeleton from '@/components/organisms/RequestTabSkeleton.vue';
import SessionTabSkeleton from '@/components/organisms/SessionTabSkeleton.vue';
import { useGetRequestsQuery } from '@/composables/queries/getRequestsQuery';
import { useGetSessionsQuery } from '@/composables/queries/getSessionsQuery';
import { useEnvironments } from '@/composables/useEnvironments';
import { ref } from 'vue';

const environment = useEnvironments();
const sessionsQuery = useGetSessionsQuery();
const requestsQuery = useGetRequestsQuery();
const { isPending, data: envs } = environment.environments();
const activeTab = ref<'sessions' | 'requests'>('sessions');

environment.$subscribe((_, state) => {
  if (state.selectedEnvironment) {
    sessionsQuery.query(state.selectedEnvironment.id);
    requestsQuery.query(state.selectedEnvironment.id);
  }
});

const switchTab = (tab: "sessions" | "requests") => {
  activeTab.value = tab;
}
</script>

<template>
  <section class="grid grid-cols-2 gap-4">
    <article class="px-6 py-4">
      <div v-if="isPending" class="flex flex-col gap-2">
        <div v-for="i in 8" class="rounded-lg p-4 w-full h-32 bg-card animate-pulse" />
      </div>
      <div v-else class="flex flex-col gap-2">
        <EnvironmentCard @click="environment.selectEnvironment(env)" v-for="env in envs" :environment="env"
          :selected="environment.selectedEnvironment?.id == env.id" />
      </div>
    </article>
    <article v-if="environment.selectedEnvironment" class="px-6 py-4 flex flex-col gap-2">
      <ul data-uk-tab class="w-full">
        <li v-bind:class="activeTab == 'sessions' ? 'uk-active' : ''"><a class="px-4 pb-3 pt-2 cursor-pointer"
            @click="() => switchTab('sessions')">Sessions</a></li>
        <li v-bind:class="activeTab == 'requests' ? 'uk-active' : ''"><a class="px-4 pb-3 pt-2 cursor-pointer"
            @click="() => switchTab('requests')">Requests</a></li>
      </ul>

      <div v-if="activeTab == 'sessions'" class="flex flex-col gap-2 mt-4">
        <div v-if="sessionsQuery.isPending.value" class="flex flex-col gap-2">
            <SessionTabSkeleton />
        </div>
        <div v-else>
          <div v-if="(sessionsQuery.data.value?.length ?? 0) > 0">
            <SessionCard v-for="session in sessionsQuery.data.value" :session="session" />
          </div>
          <div v-else>
            <p class="text-sm text-foreground">
              No sessions found for this environment. Try to refresh the page or select another environment.
            </p>
          </div>
        </div>
      </div>
      <div v-else class="flex flex-col gap-2">
        <div v-if="requestsQuery.isPending.value" class="flex flex-col gap-2">
          <RequestTabSkeleton />
        </div>
        <div v-else>
          <div v-if="(requestsQuery.data.value?.length ?? 0) > 0">
           <RequestCard v-for="request in requestsQuery.data.value" :request="request" />
          </div>
          <div v-else>
            <p class="text-sm text-foreground">
              No requests found for this environment. Try to refresh the page or select another environment.
            </p>
          </div>
        </div>
      </div>
    </article>
  </section>
</template>
