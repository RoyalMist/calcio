import React from "react";
import {MailIcon} from '@heroicons/react/solid'
import {useQuery} from "react-query";
import {API} from "../../utils/api";
import Spinner from "../../components/spinner";

interface PlayerProps {
}

interface Player {
    name: string;
    isAdmin: boolean;
    mail: string;
    avatarUrl: string;
}

function Players(props: PlayerProps) {
    const {isLoading, data} = useQuery('users', (async () => {
        const {data} = await API.get("/users")
        console.log(data);
        return data;
    }));

    if (isLoading) {
        return <Spinner loading={isLoading}/>
    } else {
        return (
            <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                {data.map((player: Player) => (
                    <li
                        key={player.mail}
                        className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow divide-y divide-gray-200"
                    >
                        <div className="flex-1 flex flex-col p-8">
                            <img className="w-32 h-32 flex-shrink-0 mx-auto rounded-full" src={player.avatarUrl} alt=""/>
                            <h3 className="mt-6 text-gray-900 text-sm font-medium">{player.name}</h3>
                            <dt className="sr-only">Role</dt>
                            {player.isAdmin && (<dd className="mt-3">
                                <span className="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full">
                                    Admin
                                </span>
                            </dd>)}
                        </div>
                        <div>
                            <div className="-mt-px flex divide-x divide-gray-200">
                                <div className="w-0 flex-1 flex">
                                    <a
                                        href={`mailto:${player.mail}`}
                                        className="group relative -mr-px w-0 flex-1 inline-flex items-center justify-center py-4 text-sm text-gray-700 font-medium border border-transparent rounded-b-lg hover:text-white hover:bg-blue-800"
                                    >
                                        <MailIcon className="w-5 h-5 text-gray-400 group-hover:text-white" aria-hidden="true"/>
                                        <span className="ml-3">Email</span>
                                    </a>
                                </div>
                            </div>
                        </div>
                    </li>
                ))}
            </ul>
        );
    }
}

export default Players;
