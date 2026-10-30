#!/usr/bin/env python3
"""Generate the additive v2 JSON contract from DevCycle's tagged public DTOs.

This is a build-time source dependency only; products share no runtime packages.
"""
from pathlib import Path
import argparse,json,re
ROOT=Path(__file__).resolve().parents[1]
parser=argparse.ArgumentParser()
parser.add_argument('--source',type=Path,default=ROOT.parent/'ai-dev-cycle-manager/internal/team/model.go',help='Exact checked-out DevCycle model.go source')
parser.add_argument('--check',action='store_true',help='Fail on generated schema drift without writing')
args=parser.parse_args()
source=args.source.read_text(encoding='utf-8')
def emit(path,value):
    text=json.dumps(value,indent=2)+'\n'
    if args.check:
        if not path.is_file() or path.read_text(encoding='utf-8')!=text:raise SystemExit('Schema drift: '+str(path))
    else:path.write_text(text,encoding='utf-8')
defs={}
def convert(t):
    if t.startswith('[]'):return {'type':['array','null'],'items':convert(t[2:])}
    if t.startswith('*'):
        v=convert(t[1:]);return {'anyOf':[v,{'type':'null'}]}
    if t in {'int','int64'}:return {'type':'integer'}
    if t=='float64':return {'type':'number'}
    if t=='bool':return {'type':'boolean'}
    if t=='string':return {'type':'string'}
    if t=='time.Time':return {'type':'string','format':'date-time'}
    return {'$ref':'#/$defs/'+t}
for name,body in re.findall(r'type (\w+) struct \{(.*?)\n\}',source,re.S):
    properties={};required=[]
    for _,typ,tag in re.findall(r'^\s*(\w+)\s+([\w.\[\]*]+)\s+`json:"([^"]+)"`',body,re.M):
        key=tag.split(',')[0];properties[key]=convert(typ)
        if 'omitempty' not in tag:required.append(key)
    if properties:defs[name]={'type':'object','additionalProperties':False,'properties':properties,'required':required}
defs['Mission']['properties']['spec']={'const':'devcycle.ai-team/v2alpha1'}
defs['Mission']['properties']['risk']={'type':'integer','minimum':0,'maximum':4}
# Additive configuration fields default when older alpha snapshots omit them.
for key in ['worktreeRoot','branchPrefix']:
    defs['Mission']['required'].remove(key)
defs['Policy']['required'].remove('requireDifferentReviewModel')
defs['Task']['properties']['role']={'enum':['lead','frontend','explorer','qa','review']}
defs['Task']['properties']['state']={'enum':['pending','blocked','ready','running','completed','failed','retry','escalated']}
defs['Mission']['properties']['state']={'enum':['pending','running','completed','failed','retry','escalated']}
schema={'$schema':'https://json-schema.org/draft/2020-12/schema','$id':'https://lifecycle-spec.local/schemas/v2/ai-team-mission.json','title':'AI Team mission v2alpha1','$ref':'#/$defs/Mission','$defs':defs}
target=ROOT/'schemas/v2';target.mkdir(exist_ok=True)
emit(target/'ai-team-mission.json',schema)
status={'$schema':schema['$schema'],'$id':'https://lifecycle-spec.local/schemas/v2/project-status.json','type':'object','additionalProperties':False,'required':['spec','projectId','state','links'],'properties':{'spec':{'const':'ai-devops.project-status/v2alpha1'},'projectId':{'type':'string','minLength':1},'state':{'enum':['development','ready','blocked','degraded','unknown']},'links':{'type':'object','additionalProperties':{'type':'string','format':'uri'}}}}
emit(target/'project-status.json',status)
print('Generated v2 contracts; legacy schemas unchanged')
