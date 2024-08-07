import type { UserModel } from './models/db/user.model';

export const UserRoles = {
	Admin: 'admin',
	Guest: 'guest',
	User: 'user'
} as const;

export type UserAccess = 'create_menu' | 'read_menu' | 'update_menu' | 'delete_menu';

export type AccessRouteMethod = 'get' | 'post' | 'put' | 'delete';
export type AccessRoute = {
	roles: ('admin' | 'guest' | 'user')[];
	access: {
		[key in UserAccess]?: 0 | 1;
	};
};

export const accessRoutes: {
	[key in AccessRouteMethod]: {
		[key: string]: AccessRoute;
	};
} = {
	get: {
		'/': {
			roles: [UserRoles.User],
			access: {}
		},
		'/about': {
			roles: [UserRoles.User],
			access: {}
		},
		'/auth/signin': {
			roles: [UserRoles.Guest],
			access: {}
		},
		'/auth/signup': {
			roles: [UserRoles.Guest],
			access: {}
		},
		'/dbg': {
			roles: [UserRoles.Guest],
			access: {}
		},
		'/dbg/col': {
			roles: [UserRoles.Guest],
			access: {}
		}
	},
	post: {},
	put: {},
	delete: {}
};

class UserAccessController {
	authorize(user: undefined | UserModel, url: string, method: string): Error | undefined {
		const route = this._pick_route(method, url);
		if (route === undefined) return new Error('undefined');
		// allow all guest routes or allow is user role is admin or allow is user has route role
		if (
			route.roles.find(
				(e) => e === UserRoles.Guest || user?.role === UserRoles.Admin || e === user?.role
			)
		)
			return;
		// if (this._has_access_to_route(user?.access ?? [], route)) return;
		return new Error('unauthorized');
	}

	_has_access_to_route(accesses: string[], route: AccessRoute): boolean {
		for (let i = 0; i < accesses.length; i++) {
			if (route.access[accesses[i] as keyof AccessRoute['access']]) {
				return true;
			}
		}
		return false;
	}

	_pick_route(method: string, url: string): AccessRoute | undefined {
		switch (method.toLowerCase()) {
			case 'get':
				return accessRoutes['get'][url];
			case 'post':
				return accessRoutes['post'][url];
			case 'put':
				return accessRoutes['put'][url];
			case 'delete':
				return accessRoutes['delete'][url];
			default:
				return;
		}
	}
}

export const uacController = new UserAccessController();
