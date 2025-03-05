<script setup lang="ts">
import EnvironmentCard from '@/components/EnvironmentCard.vue';
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
          <div v-for="i in 8" :key="i" class="w-full rounded-lg p-4 h-32 bg-card animate-pulse" />
        </div>
        <div v-else>
          <div v-if="(sessionsQuery.data.value?.length ?? 0) > 0">
            <div v-for="session in sessionsQuery.data.value" class="uk-card uk-card-body max-w-sm">
              <div class="flex flex-row gap-2 border-b border-border pb-2">
                <div class="size-5">
                  <uk-icon icon="hash"></uk-icon>
                </div>
                <h3 class="uk-card-title" v-text="session.id"></h3>
              </div>

              <div class="flex flex-row gap-2 my-4 items-center">
                <div class="size-5">
                  <uk-icon icon="user"></uk-icon>
                </div>
                <p v-text="session.userId"></p>
              </div>
              <p class="mt-4 flex flex-row gap-2">
                <span v-for="tag in session.tags" class="uk-badge uk-badge-secondary" v-text="tag"></span>
              </p>
            </div>
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
          <div v-for="i in 8" :key="i" class="w-full rounded-lg p-4 h-32 bg-card animate-pulse" />
        </div>
        <div v-else>
          <div v-if="(requestsQuery.data.value?.length ?? 0) > 0">
            <div v-for="request in requestsQuery.data.value" class="uk-card uk-card-body max-w-sm">
              <div class="flex flex-row gap-2 border-b border-border pb-2 justify-between">
                <h3 class="uk-card-title" v-text="request.name"></h3>
                <div class="flex flex-row gap-2 items-center">
                  <p class="text-sm text-foreground font-bold" v-text="request.path"></p>
                  <span class="uk-badge font-bold text-primary" v-text="request.method"></span>
                </div>
              </div>

              <div class="flex flex-col">
                <p v-text="request.headers"></p>
                <p v-text="request.body"></p>
              </div>

              <div v-if="request.userId" class="flex flex-row gap-2 my-4 items-center">
                <div class="size-5">
                  <uk-icon icon="user"></uk-icon>
                </div>
                <p v-text="request.userId"></p>
              </div>
              <div>

              </div>
            </div>
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
