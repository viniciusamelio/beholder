<script setup lang="ts">
import type { Environment } from '@/types';

const { environment, selected } = defineProps<{ environment: Environment, selected: boolean }>();

</script>

<template>
    <div class="cursor-pointer uk-card uk-card-body hover:animate-pulse max-w"
        :class="selected ? 'text-primary border-primary border-4 ring-4 ring-primary' : ''">
        <h3 class="uk-card-title flex flex-row gap-2 items-center">
            <div class="size-4">
                <uk-icon icon="chevron-right"></uk-icon>
            </div>
            <p v-text="environment.name"></p>
            <p class="text-foreground text-sm opacity-70" v-text="environment.base_url"></p>
        </h3>
        <p class="mt-4 flex flex-row gap-2">
            <span v-for="tag in environment.tags" class="uk-badge p-3" v-text="tag"></span>
        </p>
    </div>
</template>