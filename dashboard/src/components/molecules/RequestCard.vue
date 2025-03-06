<script setup lang="ts">
import type { Request } from '@/schema/schema_pb';
import RequestMethodBadge from './RequestMethodBadge.vue';

const { request } = defineProps<{ request: Request }>();
</script>

<template>
    <div class="uk-card uk-card-body w-full">
        <h3 class="uk-card-title" v-text="request.name"></h3>
        <hr class="uk-hr my-4" />
        <div class="flex flex-row gap-2 items-center">
            <RequestMethodBadge :method="request.method" />
            <span className="font-mono text-sm text-gray-300 truncate" v-text="request.path"></span>
        </div>

        <div class="flex flex-col">
            <ul class="uk-accordion" data-uk-accordion="multiple: true">
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
                    <a class="uk-accordion-title" href="#">
                        Body
                        <span class="uk-accordion-icon">
                            <uk-icon icon="chevron-down"></uk-icon>
                        </span>
                    </a>
                    <div class="uk-accordion-content">
                        <p class="text-foreground" v-text="request.body"></p>
                    </div>
                </li>
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