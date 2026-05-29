<template>
  <div class="bg-surface text-on-surface font-body-md overflow-hidden min-h-screen">
    <header class="fixed top-0 left-0 w-full z-50 px-margin-mobile md:px-margin-desktop flex items-center justify-between h-16 bg-surface/80 border-b border-white/10 backdrop-blur-md shadow-sm">
      <div class="flex items-center gap-md">
        <span class="font-headline-md text-headline-md-mobile md:text-headline-md font-bold text-primary tracking-tight">System Monitor</span>
        <div class="hidden md:block h-6 w-[1px] bg-white/10 mx-2"></div>
        <div class="hidden md:flex items-center gap-sm">
          <span class="w-2 h-2 rounded-full bg-primary status-pulse"></span>
          <span class="font-label-caps text-label-caps text-on-surface-variant uppercase">Network: Online</span>
        </div>
      </div>
      <div class="flex items-center gap-md">
        <div class="hidden md:flex flex-col items-end mr-sm">
          <span class="font-title-sm text-[12px] text-on-surface">{{ authStore.user?.name || 'Administrator' }}</span>
          <span class="font-label-caps text-[9px] text-primary">{{ authStore.user?.email || 'admin@go-iot.com' }}</span>
        </div>
        <button @click="handleLogout" class="flex items-center gap-sm p-2 hover:bg-error/10 text-on-surface-variant hover:text-error rounded-lg transition-all active:scale-95" title="Sign Out">
          <span class="material-symbols-outlined">logout</span>
        </button>
        <div class="w-8 h-8 md:w-8 md:h-8 rounded-full bg-primary-container flex items-center justify-center text-on-primary-container font-bold text-xs overflow-hidden border border-primary/20">
          <img alt="Admin" class="w-full h-full object-cover" src="../assets/images/admin_desktop_2.jpg"/>
        </div>
      </div>
    </header>

    <div class="flex h-screen pt-16">
      <aside class="hidden md:flex flex-col w-64 bg-surface-container border-r border-white/5 p-md gap-sm">
        <nav class="flex flex-col gap-xs">
          <router-link to="/" active-class="bg-primary-container text-on-primary-container shadow-sm shadow-primary/20" class="flex items-center gap-md p-md rounded-lg text-on-surface-variant hover:bg-surface-container-high transition-all">
            <span class="material-symbols-outlined">dashboard</span>
            <span class="font-label-caps text-label-caps">Overview</span>
          </router-link>
          <router-link to="/devices" active-class="bg-primary-container text-on-primary-container shadow-sm shadow-primary/20" class="flex items-center gap-md p-md rounded-lg text-on-surface-variant hover:bg-surface-container-high transition-all">
            <span class="material-symbols-outlined">developer_board</span>
            <span class="font-label-caps text-label-caps">Devices</span>
          </router-link>
          <router-link to="/admins" active-class="bg-primary-container text-on-primary-container shadow-sm shadow-primary/20" class="flex items-center gap-md p-md rounded-lg text-on-surface-variant hover:bg-surface-container-high transition-all">
            <span class="material-symbols-outlined">admin_panel_settings</span>
            <span class="font-label-caps text-label-caps">Admins</span>
          </router-link>
          <router-link to="/sectors" active-class="bg-primary-container text-on-primary-container shadow-sm shadow-primary/20" class="flex items-center gap-md p-md rounded-lg text-on-surface-variant hover:bg-surface-container-high transition-all">
            <span class="material-symbols-outlined">grid_view</span>
            <span class="font-label-caps text-label-caps">Sectors</span>
          </router-link>
          <a href="#" class="flex items-center gap-md p-md rounded-lg text-on-surface-variant hover:bg-surface-container-high transition-all">
            <span class="material-symbols-outlined">receipt_long</span>
            <span class="font-label-caps text-label-caps">Logs</span>
          </a>
        </nav>
        <div class="mt-auto p-md glass-card rounded-xl">
          <div class="flex justify-between items-center mb-sm">
            <span class="font-label-caps text-[10px] text-on-surface-variant uppercase">CPU Usage</span>
            <span class="font-data-md text-primary">24%</span>
          </div>
          <div class="w-full h-1 bg-surface-variant rounded-full overflow-hidden">
            <div class="h-full bg-primary w-[24%]"></div>
          </div>
        </div>
      </aside>

      <main class="flex-1 overflow-y-auto bg-surface-dim p-margin-mobile md:p-margin-desktop custom-scrollbar pb-32 md:pb-8">
        <router-view />
      </main>
    </div>

    <!-- Mobile Bottom Nav -->
    <nav class="md:hidden fixed bottom-0 left-0 w-full z-50 flex justify-around items-center h-20 bg-surface-container/90 border-t border-white/10 backdrop-blur-lg rounded-t-xl pb-safe">
      <router-link to="/" active-class="text-on-primary-container bg-primary-container rounded-full px-4 py-1" class="flex flex-col items-center justify-center text-on-surface-variant transition-transform active:scale-90 flex-1 py-2 mx-1 rounded-full">
        <span class="material-symbols-outlined">dashboard</span>
        <span class="font-label-caps text-[10px]">Overview</span>
      </router-link>
      <router-link to="/devices" active-class="text-on-primary-container bg-primary-container rounded-full px-4 py-1" class="flex flex-col items-center justify-center text-on-surface-variant transition-transform active:scale-90 flex-1 py-2 mx-1 rounded-full">
        <span class="material-symbols-outlined">developer_board</span>
        <span class="font-label-caps text-[10px]">Devices</span>
      </router-link>
      <router-link to="/admins" active-class="text-on-primary-container bg-primary-container rounded-full px-4 py-1" class="flex flex-col items-center justify-center text-on-surface-variant transition-transform active:scale-90 flex-1 py-2 mx-1 rounded-full">
        <span class="material-symbols-outlined">admin_panel_settings</span>
        <span class="font-label-caps text-[10px]">Admins</span>
      </router-link>
      <router-link to="/sectors" active-class="text-on-primary-container bg-primary-container rounded-full px-4 py-1" class="flex flex-col items-center justify-center text-on-surface-variant transition-transform active:scale-90 flex-1 py-2 mx-1 rounded-full">
        <span class="material-symbols-outlined">grid_view</span>
        <span class="font-label-caps text-[10px]">Sectors</span>
      </router-link>
    </nav>
  </div>
</template>

<script setup>
import { useAuthStore } from '../store/auth'
import { useRouter } from 'vue-router'

const authStore = useAuthStore()
const router = useRouter()

const handleLogout = () => {
    authStore.logout()
    router.push('/login')
}
</script>
