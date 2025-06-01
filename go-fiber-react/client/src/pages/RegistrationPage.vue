<script setup lang="ts">
import useRegistrationPage from './RegistrationPage';

const { email, password, username, handleSubmit, isPwd } = useRegistrationPage();
</script>

<template>
  <div>
    <q-label> Username {{ username }}</q-label>
    <q-input v-model="username" filled />
  </div>

  <div>
    <q-label> Email </q-label>
    <q-input v-model="email" filled />
  </div>
  <q-input
    v-model="password"
    filled
    :type="isPwd ? 'password' : 'text'"
    hint="Password with toggle"
  >
    <template v-slot:append>
      <q-icon
        :name="isPwd ? 'visibility_off' : 'visibility'"
        class="cursor-pointer"
        @click="isPwd = !isPwd"
      />
    </template>
  </q-input>

  <q-btn label="submit" icon="send" class="bg-blue" @click="handleSubmit" />
</template>
