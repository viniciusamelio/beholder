<script setup lang="ts">
import type { Session } from '@/schema/schema_pb';

const { session, selected } = defineProps<{ session: Session, selected : boolean }>();
</script>

<template>
    <div class="uk-card uk-card-body max-w-sm cursor-pointer" :class="selected ? 'border-primary border-4 ring-4 ring-primary' : ''">
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
</template>