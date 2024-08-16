const strictLoad = (env: string): string => {
  const val = import.meta.env[env];
  if (!val) throw new Error(`${env} environment variable is undefined!`);
  return val;
};

export const env = {
  dev: import.meta.env.DEV,
  ssr: import.meta.env.SSR,
  mode: import.meta.env.MODE,
  prod: import.meta.env.PROD,
  baseUrl: import.meta.env.BASE_URL,
  app: {
    test: strictLoad("VITE_TEST"),
  },
  backend: {
    baseUrl: import.meta.env.DEV ? "localhost:3000" : window.location.host,
  },
};
