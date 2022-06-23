// App imports
import { TOGGLE_APP_STATUS, TOGGLE_NAV, TOGGLE_WC_CC_CATEGORY } from '../actions/Ui.actions';
import { Action, DynamicCategoryToggleAction } from '../../shared/types/types';
import { ReducerDescriptor } from '../../shared/utilities/Reducer.utils';

export const STORE_SECTION_UI = 'ui';

interface UIState {
    isDeployInProgress: boolean;
    navExpanded: boolean;
    wcCcCategoryExpanded: { [category: string]: boolean };
}

function uiReducer(state: UIState, action: Action) {
    const newState = { ...state };
    switch (action.type) {
        case TOGGLE_APP_STATUS:
            newState['isDeployInProgress'] = !state.isDeployInProgress;
            console.log('APP UI:', newState);
            break;
        case TOGGLE_NAV:
            newState['navExpanded'] = !state.navExpanded;
            console.log('APP UI:', newState);
            break;
        case TOGGLE_WC_CC_CATEGORY:
            newState['wcCcCategoryExpanded'] = createStateToggleCategory(state.wcCcCategoryExpanded, action as DynamicCategoryToggleAction);
            console.log('APP UI:', newState);
            break;
    }
    return newState;
}

export const uiReducerDescriptor = {
    name: 'generic ui reducer',
    reducer: uiReducer,
    actionTypes: [TOGGLE_APP_STATUS, TOGGLE_NAV, TOGGLE_WC_CC_CATEGORY],
    storeSection: STORE_SECTION_UI,
} as ReducerDescriptor;

// given an old categoryExpanded object, create a new categoryExpanded object (with the category toggled)
function createStateToggleCategory(oldCategoryObject: any, action: DynamicCategoryToggleAction): any {
    const oldToggleValue = oldCategoryObject[action.category] || false;
    return { ...oldCategoryObject, [action.category]: !oldToggleValue };
}