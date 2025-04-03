<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
import RequestCard from '@/components/molecules/RequestCard.vue';
import RequestPill from '@/components/molecules/RequestPill.vue';
import SessionCard from '@/components/molecules/SessionCard.vue';
import RequestTabSkeleton from '@/components/organisms/RequestTabSkeleton.vue';
import SessionTabSkeleton from '@/components/organisms/SessionTabSkeleton.vue';
import { useGetRequestsQuery } from '@/composables/queries/getRequestsQuery';
import { useGetSessionQuery } from '@/composables/queries/getSessionQuery';
import { useGetSessionsQuery } from '@/composables/queries/getSessionsQuery';
import { useEnvironments } from '@/composables/useEnvironments';
import { ref, watch } from 'vue';

const environment = useEnvironments();
const sessionsQuery = useGetSessionsQuery();
const requestsQuery = useGetRequestsQuery();
const sessionQuery = useGetSessionQuery();
const { isPending, data: envs } = environment.environments();
const activeTab = ref<'sessions' | 'requests'>('sessions');
const selectedSessionId = ref<number | null>(null);

environment.$subscribe((_, state) => {
  if (state.selectedEnvironment) {
    sessionsQuery.query(state.selectedEnvironment.id);
    requestsQuery.query(state.selectedEnvironment.id);
  }
});
watch(selectedSessionId, () => {
  if (!selectedSessionId.value) return;
  sessionQuery.query(selectedSessionId.value.toString());
});

const switchTab = (tab: "sessions" | "requests") => {
  activeTab.value = tab;
}
</script>

<template>
  <section class="grid grid-cols-2 gap-4">
    <article v-if="!selectedSessionId" class="px-6 py-4">
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
            <SessionCard
              @click="selectedSessionId != session.id ? selectedSessionId = session.id : selectedSessionId = null"
              :selected="session.id == selectedSessionId" v-for="session in sessionsQuery.data.value"
              :session="session" />
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
          <div class="flex flex-col gap-2" v-if="(requestsQuery.data.value?.length ?? 0) > 0">
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
    <article v-if="selectedSessionId" class="px-6 py-4 flex flex-col gap-2 border-l-2 border-border min-h-screen">
      <div v-if="sessionQuery.isPending.value">
        <SessionTabSkeleton />
      </div>
      <div v-else-if="sessionQuery.isError.value">
        <p class="text-sm text-foreground">
          No session found for this environment. Try to refresh the page or select another environment.
        </p>
      </div>
      <div v-else-if="sessionQuery.data.value">
        <div class="flex flex-col gap-2">
          <div class="flex flex-row gap-2 justify-between items-center mb-4">
            <button class="uk-button uk-button-primary" @click="selectedSessionId = null" v-if="selectedSessionId">
              <uk-icon icon="chevron-left" class="text-primary" height="24" width="24" />
            </button>
            <button class="uk-btn uk-btn-primary flex flex-row gap-2 items-center justify-center">
              <uk-icon icon="rotate-ccw" />
              Replay
            </button>
          </div>
          
          <RequestCard v-for="request in sessionQuery.data.value.requests" :request="request" />
        </div>
      </div>
    </article>
  </section>
</template>
