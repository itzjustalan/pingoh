import { browser } from '$app/environment';
import { writable } from 'svelte/store';

type AppTheme = 'light' | 'dark';

function createStore() {
	const darkPrefered =
		browser && window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
	const { subscribe, set, update } = writable<AppTheme>(darkPrefered ? 'dark' : 'light');

	return {
		set,
		subscribe,
		dark: () => set('dark'),
		light: () => set('light'),
		toggle: () => update((theme) => (theme === 'dark' ? 'light' : 'dark'))
	};
}

export const appTheme = createStore();

if (browser) {
	window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (event) => {
		appTheme.set(event.matches ? 'dark' : 'light');
	});
}

appTheme.subscribe((theme) => {
	if (!browser) return;
	if (theme === 'dark') document.body.classList.add('dark-mode');
	else document.body.classList.remove('dark-mode');
});
