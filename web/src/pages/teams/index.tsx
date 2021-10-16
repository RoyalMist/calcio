import React, {Fragment, useState} from "react";
import SectionHeader from "../../components/SectionHeader";
import {useAuth} from "../../stores/authentication";
import Spinner from "../../components/Spinner";
import SlideOver from "../../components/SlideOver";
import {ent_Team, ent_User, TeamsService, UsersService} from "../../gen";
import {useMutation, useQuery, useQueryClient} from "react-query";
import {Listbox, Transition} from "@headlessui/react";
import {CheckIcon, SelectorIcon} from "@heroicons/react/solid";
import {classNames} from "../../utils/classes";
import {SubmitHandler, useForm} from "react-hook-form";


const Teams = () => {
    const {token, isAdmin, paseto} = useAuth();
    const [open, setOpen] = useState(false);
    const [selectedTeammate, setSelectedTeammate] = useState<ent_User>({name: "No"});
    const {handleSubmit} = useForm<ent_User>();
    const onSubmit: SubmitHandler<ent_User> = () => createTeam.mutate(selectedTeammate);
    const newTeam = () => {
        setSelectedTeammate({name: "No"});
        setOpen(true);
    }

    const queryClient = useQueryClient();
    const usersQuery = useQuery('users', async () => {
        const data = await UsersService.getUsersService(token);
        return data.filter(user => user.id !== paseto?.user_id);
    });
    const teamsQuery = useQuery('teams', async () => TeamsService.getTeamsService(token))
    const createTeam = useMutation('teams', async (values: ent_User) => {
        return await TeamsService.putTeamsService(token, values.id)
    }, {
        onSuccess: async () => {
            setOpen(false);
            return await queryClient.invalidateQueries('teams');
        }
    });

    const avatar = (name: string | undefined): string => {
        if (name === undefined) {
            return "";
        }

        const splited = name.split(" & ");
        if (splited.length > 1) {
            return `${splited[0].substr(0, 1).toUpperCase()}${splited[1].substr(0, 1).toUpperCase()}`;
        } else {
            return name.substr(0, 2).toUpperCase();
        }
    }

    return (
        <>
            <Spinner loading={usersQuery.isLoading || teamsQuery.isLoading || createTeam.isLoading}/>
            <SectionHeader action={newTeam}>{isAdmin() ? "Teams" : "My Teams"}</SectionHeader>
            <SlideOver open={open} close={() => setOpen(false)} title="Create New Team">
                <form className="space-y-6" onSubmit={handleSubmit(onSubmit)}>
                    <Listbox value={selectedTeammate} onChange={setSelectedTeammate}>
                        <Listbox.Label className="block text-sm font-medium text-gray-700">With a teammate?</Listbox.Label>
                        <div className="mt-1 relative">
                            <Listbox.Button
                                className="relative w-full bg-white border border-gray-300 rounded-md shadow-sm pl-3 pr-10 py-2 text-left cursor-default focus:outline-none focus:ring-1 focus:ring-blue-500 focus:border-blue-500 sm:text-sm">
                                    <span className="flex items-center">
                                        <span className="ml-3 block truncate">{selectedTeammate.name}</span>
                                    </span>
                                <span className="ml-3 absolute inset-y-0 right-0 flex items-center pr-2 pointer-events-none">
                                        <SelectorIcon className="h-5 w-5 text-gray-400" aria-hidden="true"/>
                                    </span>
                            </Listbox.Button>
                            <Transition as={Fragment} leave="transition ease-in duration-100" leaveFrom="opacity-100" leaveTo="opacity-0">
                                <Listbox.Options
                                    className="absolute z-10 mt-1 w-full bg-white shadow-lg max-h-56 rounded-md py-1 text-base ring-1 ring-black ring-opacity-5 overflow-auto focus:outline-none sm:text-sm">
                                    {usersQuery.isSuccess && usersQuery.data.map((user) => (
                                        <Listbox.Option
                                            key={user.id}
                                            className={({active}) =>
                                                classNames(
                                                    active ? 'text-white bg-blue-600' : 'text-gray-900',
                                                    'cursor-default select-none relative py-2 pl-3 pr-9'
                                                )
                                            }
                                            value={user}
                                        >
                                            {({selected, active}) => (
                                                <>
                                                    <div className="flex items-center">
                                                            <span className={classNames(selected ? 'font-semibold' : 'font-normal', 'ml-3 block truncate')}>
                                                                {user.name}
                                                            </span>
                                                    </div>
                                                    {selected ? (
                                                        <span
                                                            className={classNames(
                                                                active ? 'text-white' : 'text-blue-600',
                                                                'absolute inset-y-0 right-0 flex items-center pr-4'
                                                            )}
                                                        >
                                                                <CheckIcon className="h-5 w-5" aria-hidden="true"/>
                                                            </span>
                                                    ) : null}
                                                </>
                                            )}
                                        </Listbox.Option>
                                    ))}
                                </Listbox.Options>
                            </Transition>
                        </div>
                    </Listbox>
                    <div>
                        <button type="submit" className="w-full btn btn-primary">
                            Save
                        </button>
                    </div>
                </form>
            </SlideOver>
            <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {teamsQuery.isSuccess && teamsQuery.data.map((team: ent_Team) => (
                    <li
                        key={team.id}
                        className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow divide-y divide-gray-200"
                    >
                        <div className="flex-1 flex flex-col p-8">
                            <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                                <span className="text-xl font-medium leading-none text-white">{avatar(team.name)}</span>
                            </span>
                            <h3 className="mt-6 text-gray-900 text-sm font-medium">{team.name}</h3>
                        </div>
                    </li>
                ))}
            </ul>
        </>
    );
};

export default Teams;
