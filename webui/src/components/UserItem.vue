<script>
export default {
    props: ["userData"]
}
</script>

<template>
    <a :href="`#/${userData.username}`">
        {{ userData.username }}
    </a>
</template>

<style></style>