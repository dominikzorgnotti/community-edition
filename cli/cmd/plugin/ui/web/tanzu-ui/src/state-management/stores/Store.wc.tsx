// React imports
import React, { createContext, ReactNode, useReducer } from 'react';

// App imports
import { StoreDispatch } from '../../shared/types/types';
import wcReducer from '../reducers/Wizard.reducer';

const initialState = {
    data: {
        ccAttributes: {},
        SELECTED_MANAGEMENT_CLUSTER: '',
        SELECTED_CLUSTER_CLASS: '',
        AVAILABLE_CLUSTER_CLASSES: [],
    },
    ui: {
        wcCcCategoryExpanded: {},
    },
};

const WcStore = createContext<{
    state: { [key: string]: any };
    dispatch: StoreDispatch;
}>({
    state: initialState,
    dispatch: () => null,
});

const WcProvider: React.FC<{ children: ReactNode }> = ({ children }: { children: ReactNode }) => {
    const [state, dispatch] = useReducer(wcReducer, initialState);

    return <WcStore.Provider value={{ state, dispatch }}>{children}</WcStore.Provider>;
};

export { WcStore, WcProvider };