import React, {useState} from "react";
import {PencilAltIcon, TrashIcon} from '@heroicons/react/solid'
import Spinner from "../../components/Spinner";
import SectionHeader from "../../components/SectionHeader";
import {api_user, ent_User, UsersService} from "../../gen";
import SlideOver from "../../components/SlideOver";
import {useMutation, useQuery, useQueryClient} from "react-query";
import {useAuth} from "../../stores/authentication";

function Users() {
    const {token} = useAuth();
    const [open, setOpen] = useState(false);
    const [selectedUser, setSelectedUser] = useState<api_user>({name: "", password: "", admin: false});

    const newUser = () => {
        setSelectedUser({name: "", password: "", admin: false});
        setOpen(true);
    }

    const editUser = (user: api_user) => {
        setSelectedUser(user);
        setOpen(true);
    }

    const queryClient = useQueryClient();
    const usersQuery = useQuery('users', async () => UsersService.getUsersService(token));
    const saveUser = useMutation('users', async (values: api_user) => {
        if (values.id != undefined) {
            return await UsersService.putUsersService(token, values);
        } else {
            return await UsersService.postUsersService(token, values);
        }
    }, {
        onSuccess: async () => {
            setOpen(false);
            return await queryClient.invalidateQueries('users');
        }
    });

    const deleteUser = useMutation('users', async (id: string) => {
        return await UsersService.deleteUsersService(token, id)
    }, {
        onSuccess: async () => {
            return await queryClient.invalidateQueries('users');
        }
    })

    return (
        <>
            <Spinner loading={usersQuery.isLoading || saveUser.isLoading}/>
            <SectionHeader action={newUser}>Users</SectionHeader>
            <SlideOver open={open} close={() => setOpen(false)} title="Edit User">
                <Formik
                    initialValues={selectedUser}
                    onSubmit={(values) => {
                        saveUser.mutate({...values});
                    }}>
                    <Form className="space-y-6">
                        {selectedUser.id === undefined &&
                        <SimpleField
                            type="text"
                            name="name"
                            placeholder="John"
                        >
                            Name
                        </SimpleField>
                        }
                        <SimpleField
                            type="password"
                            name="password"
                            placeholder="**********"
                        >
                            Password
                        </SimpleField>
                        <div>
                            <button type="submit" className="w-full btn btn-primary">
                                Save
                            </button>
                        </div>
                    </Form>
                </Formik>
            </SlideOver>
            <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {usersQuery.isSuccess && usersQuery.data.map((user: ent_User) => (
                    <li
                        key={user.id}
                        className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow divide-y divide-gray-200"
                    >
                        <div className="flex-1 flex flex-col p-8">
                                <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                                    <span className="text-xl font-medium leading-none text-white">{user.name?.substr(0, 2).toUpperCase()}</span>
                                </span>
                            <h3 className="mt-6 text-gray-900 text-sm font-medium">{user.name}</h3>
                            <dt className="sr-only">Role</dt>
                            {user.admin && (<dd className="mt-3">
                                <span className="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full">
                                    Admin
                                </span>
                            </dd>)}
                        </div>
                        <div>
                            <div className="-mt-px flex divide-x divide-gray-200">
                                <div className="w-0 flex-1 flex">
                                    <button
                                        onClick={() => editUser(user)}
                                        className="group relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm text-gray-700 font-medium border border-transparent rounded-bl-lg hover:text-white hover:bg-blue-800">
                                        <PencilAltIcon className="w-5 h-5 text-gray-400 group-hover:text-white" aria-hidden="true"/>
                                        <span className="ml-3">Edit</span>
                                    </button>
                                    <button
                                        onClick={() => deleteUser.mutate(user.id != undefined ? user.id : "")}
                                        className="group relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm text-gray-700 font-medium border border-transparent rounded-br-lg hover:text-white hover:bg-red-800">
                                        <TrashIcon className="w-5 h-5 text-gray-400 group-hover:text-white" aria-hidden="true"/>
                                        <span className="ml-3">Delete</span>
                                    </button>
                                </div>
                            </div>
                        </div>
                    </li>
                ))}
            </ul>
        </>
    );
}

export default Users;
