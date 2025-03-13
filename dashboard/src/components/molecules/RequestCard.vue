<script setup lang="ts">
import type { Request } from '@/schema/schema_pb';
import RequestMethodBadge from './RequestMethodBadge.vue';
import { ref } from 'vue';

const { request } = defineProps<{ request: Request }>();

const activeTab = ref<"request"|"response">("request");

</script>

<template>
    <div class="uk-card uk-card-body w-full p-0 rounded-lg">
        <h3 class="uk-card-title pt-5 p-2 pb-0 flex items-end" v-text="request.name"></h3>
        <hr class="uk-hr mt-2 mb-0" />
        <div class="flex flex-row gap-2 items-center p-2">
            <RequestMethodBadge :method="request.method" />
            <span className="font-mono text-sm text-gray-300 truncate" v-text="request.path"></span>
        </div>
        <ul class="uk-tab-alt text-sm rounded-none">
            <li @click="activeTab = 'request'" :class="activeTab == 'request' ? 'uk-active' : ''"><a href="#">Request</a></li>
            <li @click="activeTab = 'response'" :class="activeTab == 'response' ? 'uk-active' : ''"><a href="#">Response <span class="text-xs ml-2 text-emerald-400 font-semi-bold">200</span> </a></li>
        </ul>

        <div class="flex flex-col">
            <ul v-if="activeTab == 'request'" class="uk-accordion hover:bg-muted px-2 py-0" data-uk-accordion="multiple: true">
                <li v-if="request.headers">
                    <a class="uk-accordion-title" href="#">
                        Headers
                        <span class="uk-accordion-icon">
                            <uk-icon icon="chevron-down"></uk-icon>
                        </span>
                    </a>
                    <div class="uk-accordion-content">
                        <p class="text-foreground" v-text="request.headers"></p>
                    </div>
                </li>
                <li v-else-if="request.body">
                    <a class="uk-accordion-title p-0 py-1 rounded-xl text-sm" href="#">
                        Body
                        <span class="uk-accordion-icon">
                            <uk-icon icon="chevron-down"></uk-icon>
                        </span>
                    </a>
                    <div class="uk-accordion-content mt-2">
                        <p class="text-foreground text-sm w-full" v-text="request.body"></p>
                    </div>
                </li>
            </ul>
            <ul v-else>

            </ul>
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
</template>