import React, { createContext, Reducer } from 'react';

type AuthState = {
  accessToken: string;
};

type AuthContextProps = {
  auth: AuthState;
  authenticate?: (token: string) => void;
  unauthenticate?: () => void;
};

const initialState: AuthState = {
  accessToken: '',
};

const AuthContext = createContext<AuthContextProps>({ auth: initialState });

const AUTHENTICATED = 'contexts.auth.authenticated';
const UNAUTHENTICATED = 'contexts.auth.unauthenticated';

type AuthAction =
  | {
      type: typeof AUTHENTICATED;
      payload: { accessToken: string };
    }
  | {
      type: typeof UNAUTHENTICATED;
      payload: {};
    };

const reducer: Reducer<AuthState, AuthAction> = (state, action) => {
  switch (action.type) {
    case AUTHENTICATED:
      return { ...state, accessToken: action.payload.accessToken };

    case UNAUTHENTICATED:
      return { ...state, accessToken: undefined };

    default:
      return state;
  }
};

type AuthContextProviderProps = {
  accessToken?: string;
};

export const AuthContextProvider: React.FC<AuthContextProviderProps> = ({
  children,
}) => {
  const [state, dispatch] = React.useReducer(reducer, initialState);
  return (
    <AuthContext.Provider
      value={{
        auth: { ...state },
        authenticate: (accessToken) =>
          dispatch({ type: AUTHENTICATED, payload: { accessToken } }),
        unauthenticate: () => dispatch({ type: UNAUTHENTICATED, payload: {} }),
      }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const AuthContextConsumer = AuthContext.Consumer;

export default AuthContext;
