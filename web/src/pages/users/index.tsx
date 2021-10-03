import React, {useState} from "react";
import {MailIcon} from '@heroicons/react/solid'
import Spinner from "../../components/spinner";
import SectionHeader from "../../components/section-header";
import {ent_User} from "../../gen";
import SlideOver from "../../components/slide-over";

function Users() {
    const isLoading = false;
    const [open, setOpen] = useState(false);
    const data: ent_User[] = [];


    if (isLoading) {
        return <Spinner loading={isLoading}/>
    } else {
        return (
            <>
                <SectionHeader action={() => setOpen(true)}>Users</SectionHeader>
                <SlideOver open={open} close={() => setOpen(false)} title="Edit User">
                    <h1>Hello</h1>
                </SlideOver>
                <ul role="list" className="grid grid-cols-1 gap-6 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4">
                    {data.map((player: ent_User) => (
                        <li
                            key={player.id}
                            className="col-span-1 flex flex-col text-center bg-white rounded-lg shadow divide-y divide-gray-200"
                        >
                            <div className="flex-1 flex flex-col p-8">
                                <span className="inline-flex items-center justify-center h-14 w-14 rounded-full bg-gray-500">
                                    <span className="text-xl font-medium leading-none text-white">TW</span>
                                </span>
                                <h3 className="mt-6 text-gray-900 text-sm font-medium">{player.name}</h3>
                                <dt className="sr-only">Role</dt>
                                {player.admin && (<dd className="mt-3">
                                <span className="px-2 py-1 text-green-800 text-xs font-medium bg-green-100 rounded-full">
                                    Admin
                                </span>
                                </dd>)}
                            </div>
                            <div>
                                <div className="-mt-px flex divide-x divide-gray-200">
                                    <div className="w-0 flex-1 flex">
                                        <a
                                            href={`mailto:${player.name}`}
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
            </>
        );
    }
}

export default Users;
