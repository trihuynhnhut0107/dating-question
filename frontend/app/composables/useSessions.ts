import type { CreateSessionDto, SessionDto } from "~/types/session";

const SESSIONS_ENDPOINT = "/sessions";

export const useSessions = () => {
  const createSession = (payload: CreateSessionDto) =>
    useApi<SessionDto>(SESSIONS_ENDPOINT, {
      method: "POST",
      body: payload,
    });

  return {
    createSession,
  };
};
