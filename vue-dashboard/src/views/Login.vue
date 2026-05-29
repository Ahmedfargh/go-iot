<template>
<div class="min-h-screen flex items-center justify-center p-margin-mobile md:p-margin-desktop font-body-md text-on-surface selection:bg-primary selection:text-on-primary">
  <main class="w-full max-w-[440px] animate-in fade-in zoom-in duration-700">
    <div class="flex flex-col items-center mb-xl">
      <div class="w-12 h-12 bg-primary flex items-center justify-center rounded-xl mb-md shadow-[0_0_20px_rgba(78,222,163,0.3)]">
        <span class="material-symbols-outlined text-on-primary text-[28px]">developer_board</span>
      </div>
      <h1 class="font-headline-md text-headline-md text-primary tracking-tight">System Monitor</h1>
      <p class="font-label-caps text-label-caps text-on-surface-variant mt-xs">Industrial IoT Interface v4.0.2</p>
    </div>

    <section class="glass-panel p-xl rounded-xl shadow-2xl relative overflow-hidden" ref="loginCard">
      <div class="absolute top-0 right-0 w-16 h-16 bg-gradient-to-bl from-primary/10 to-transparent"></div>
      <header class="mb-lg">
        <h2 class="font-title-sm text-title-sm text-on-surface">Terminal Authorization</h2>
        <p class="text-on-surface-variant text-body-md mt-xs">Provide credentials to access monitoring node.</p>
      </header>

      <form @submit.prevent="handleLogin" class="space-y-lg">
        <div class="space-y-xs">
          <label class="font-label-caps text-label-caps text-on-surface-variant flex justify-between" for="email">
            Operator Email
            <span class="text-primary/50 text-[10px]">REQUIRED</span>
          </label>
          <div class="relative group">
            <span class="material-symbols-outlined absolute left-md top-1/2 -translate-y-1/2 text-on-surface-variant group-focus-within:text-primary transition-colors">alternate_email</span>
            <input v-model="email" class="w-full bg-surface-container-lowest border border-white/10 rounded-lg py-3 pl-11 pr-md font-data-md text-data-md text-on-surface placeholder:text-on-surface-variant/30 focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" id="email" placeholder="admin@sector-7.io" type="email"/>
          </div>
        </div>

        <div class="space-y-xs">
          <label class="font-label-caps text-label-caps text-on-surface-variant flex justify-between" for="password">
            Security Token
            <span class="text-primary/50 text-[10px]">AES-256</span>
          </label>
          <div class="relative group">
            <span class="material-symbols-outlined absolute left-md top-1/2 -translate-y-1/2 text-on-surface-variant group-focus-within:text-primary transition-colors">lock</span>
            <input v-model="password" class="w-full bg-surface-container-lowest border border-white/10 rounded-lg py-3 pl-11 pr-md font-data-md text-data-md text-on-surface placeholder:text-on-surface-variant/30 focus:outline-none focus:border-primary focus:ring-1 focus:ring-primary transition-all" id="password" placeholder="••••••••••••" type="password"/>
          </div>
        </div>

        <div v-if="error" class="bg-error/10 border border-error/20 rounded-lg p-sm flex items-center gap-sm text-error text-body-md animate-in slide-in-from-top-2">
          <span class="material-symbols-outlined text-[18px]">error</span>
          {{ error }}
        </div>

        <div class="flex items-center justify-between py-xs">
          <label class="flex items-center gap-sm cursor-pointer group">
            <div class="relative flex items-center">
              <input class="peer appearance-none w-4 h-4 rounded-xs border border-white/20 checked:bg-primary checked:border-primary transition-all" type="checkbox"/>
              <span class="material-symbols-outlined absolute text-[12px] text-on-primary opacity-0 peer-checked:opacity-100 left-0.5">check</span>
            </div>
            <span class="text-body-md text-on-surface-variant group-hover:text-on-surface transition-colors">Persist session</span>
          </label>
          <a class="text-body-md text-secondary hover:text-secondary-fixed-dim transition-colors font-medium" href="#">Forgot Password</a>
        </div>

        <button class="w-full bg-primary-container hover:bg-primary text-on-primary-container font-title-sm text-title-sm py-4 rounded-lg flex items-center justify-center gap-sm transition-all duration-300 active:scale-[0.98] shadow-[0_4px_12px_rgba(16,185,129,0.2)]" type="submit" :disabled="isLoading">
          <span v-if="isLoading" class="material-symbols-outlined animate-spin text-[20px]">refresh</span>
          <template v-else>
            Sign In
            <span class="material-symbols-outlined">login</span>
          </template>
        </button>
      </form>

      <footer class="mt-xl pt-lg border-t border-white/5 flex flex-col items-center gap-md">
        <div class="flex gap-lg">
          <button class="flex items-center gap-xs font-label-caps text-label-caps text-on-surface-variant hover:text-primary transition-colors">
            <span class="material-symbols-outlined text-[16px]">fingerprint</span>
            Biometric
          </button>
          <button class="flex items-center gap-xs font-label-caps text-label-caps text-on-surface-variant hover:text-primary transition-colors">
            <span class="material-symbols-outlined text-[16px]">qr_code_2</span>
            Mobile App
          </button>
        </div>
      </footer>
    </section>

    <!-- System Background Information -->
    <div class="mt-xl grid grid-cols-2 gap-md opacity-40">
      <div class="glass-panel p-md rounded-lg flex flex-col gap-xs">
        <span class="font-label-caps text-[10px] text-on-surface-variant">UPLINK STATUS</span>
        <div class="flex items-center gap-xs font-data-md text-data-md text-primary">
          <span class="w-1.5 h-1.5 rounded-full bg-primary animate-pulse"></span>
          STABLE
        </div>
      </div>
      <div class="glass-panel p-md rounded-lg flex flex-col gap-xs">
        <span class="font-label-caps text-[10px] text-on-surface-variant">NODE ID</span>
        <div class="font-data-md text-data-md text-on-surface">
          PX-7721-B
        </div>
      </div>
    </div>

    <!-- System Image Decoration -->
    <div class="fixed bottom-0 left-0 w-full h-[265px] -z-10 overflow-hidden pointer-events-none opacity-20 grayscale">
      <img alt="Server background" class="w-full h-full object-cover object-center" src="../assets/images/server_bg.jpg"/>
    </div>
  </main>
</div>
</template>

<script setup>
import { onMounted, ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../store/auth';

const router = useRouter()
const authStore = useAuthStore()
const loginCard = ref(null)

const email = ref('admin@go-iot.com')
const password = ref('password123')
const error = ref('')
const isLoading = ref(false)

const handleLogin = async () => {
    error.value = ''
    isLoading.value = true
    try {
        await authStore.login(email.value, password.value)
        router.push('/')
    } catch (err) {
        error.value = err.response?.data?.message || 'Authorization failed. Check credentials.'
        console.error('Login error:', err)
    } finally {
        isLoading.value = false
    }
}

onMounted(() => {
    setTimeout(() => {
        if (loginCard.value) {
            loginCard.value.style.borderColor = 'rgba(78, 222, 163, 0.4)';
            setTimeout(() => {
                loginCard.value.style.borderColor = 'rgba(255, 255, 255, 0.08)';
            }, 1000);
        }
    }, 500);
});
</script>
